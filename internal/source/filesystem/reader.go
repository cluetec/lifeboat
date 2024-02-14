/*
 * Copyright 2023 cluetec GmbH
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filesystem

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"

	globalConfig "github.com/cluetec/lifeboat/internal/config"
)

type Reader struct {
	buffer *bytes.Buffer
}

func NewReader(rc *globalConfig.ResourceConfig) (*Reader, error) {
	c, err := newConfig(rc)
	if err != nil {
		slog.Error("error while initializing filesystem source config", "error", err)
		return nil, err
	}

	slog.Debug("filesystem source config loaded", "config", rc)

	fileInfo, err := os.Stat(c.Path)
	if err != nil {
		return nil, err
	}

	r := &Reader{
		buffer: bytes.NewBuffer(make([]byte, 0)),
	}

	if fileInfo.IsDir() {
		err = r.prepareDir(c.Path)
	} else {
		err = r.prepareFile(c.Path)
	}
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	slog.Debug("filesystem source read got called")
	return r.buffer.Read(b)
}

func (r *Reader) Close() error {
	slog.Debug("closing filesystem reader")
	return nil
}

// prepareDir reads in the source directory for the backup as a gzipped tar archive and sets
// the contents as a bytes.Buffer to the Reader so it can be written during the backup.
//
// Symlinks found in the source directory are resolved so the file the link points to will be put into the resulting backup archive.
func (r *Reader) prepareDir(srcPath string) error {
	slog.Debug("preparing filesystem directory backup")
	mw := io.MultiWriter(r.buffer)

	gw := gzip.NewWriter(mw)
	tw := tar.NewWriter(gw)

	if err := filepath.WalkDir(srcPath, writeFileIntoArchive(tw, srcPath, "")); err != nil {
		return err
	}

	if err := tw.Close(); err != nil {
		return err
	}

	if err := gw.Close(); err != nil {
		return err
	}

	return nil
}

// writeFileIntoArchive returns the function which reads the current walked file and writes it into
// the given tar writer.
// When parentDir is set the current file will be placed in the given directory inside the resulting archive. Mainly used for symlinks on
// directories so their contents can be resolved into the correct location.
func writeFileIntoArchive(tw *tar.Writer, srcPath string, parentDir string) func(path string, entry fs.DirEntry, err error) error {
	return func(path string, entry fs.DirEntry, err error) error {
		slog.Info("====== NEXT FILE ======")
		slog.Debug("paths", "srcPath", srcPath, "path", path)
		if err != nil {
			return err
		}

		// exclude root dir from resulting archive
		if path == srcPath {
			return nil
		}

		fileInfo, err := entry.Info()
		if err != nil {
			return err
		}
		slog.Debug("walking...", "fileInfo", fileInfo.Name())

		// check if entry is a symlink
		if entry.Type()&fs.ModeSymlink == fs.ModeSymlink {
			resolvedFileInfo, resolvedPath, err := resolveSymLink(srcPath, path)
			if err != nil {
				return err
			}

			fileInfo = *resolvedFileInfo

			// check if symlink is pointing to a dir
			if fileInfo.IsDir() {
				slog.Debug("symlink points to a directory - iterate over files", "resolvedPath", resolvedPath)
				if err := filepath.WalkDir(resolvedPath, writeFileIntoArchive(tw, resolvedPath, fileInfo.Name())); err != nil {
					return err
				}
			}
		}

		header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
		if err != nil {
			return err
		}

		// preserve folder structure inside tar archive, f.e. for files in nested directories
		header.Name, err = filepath.Rel(srcPath, filepath.Join(filepath.Dir(path), parentDir, fileInfo.Name()))
		if err != nil {
			return err
		}

		slog.Debug("writing header", "header", header.Name)
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// write the file contents into the archive
		if !fileInfo.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			slog.Info("writing file", "originalPath", path, "resultingFile", header.Name)
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil
	}
}

// resolveSymLink returns the fileInfo and the resolved path of the file which a soft symlinks point to
func resolveSymLink(srcPath, linkPath string) (*fs.FileInfo, string, error) {
	resolvedPath, err := os.Readlink(linkPath)
	if err != nil {
		return nil, "", err
	}

	if !filepath.IsAbs(resolvedPath) {
		resolvedPath = fmt.Sprintf("%s/%s", srcPath, resolvedPath)
	}

	fileInfo, err := os.Stat(resolvedPath)
	slog.Debug("resolved symlink", "resolvedPath", resolvedPath, "resolvedFileInfo", fileInfo.Name())
	return &fileInfo, resolvedPath, err
}

// prepareFile opens the source file with the given path, reads its contents chunkwise and writes it
// into the buffer of the reader
func (r *Reader) prepareFile(filePath string) error {
	slog.Debug("preparing filesystem file backup")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var count int
	reader := bufio.NewReader(file)
	r.buffer = bytes.NewBuffer(make([]byte, 0))
	part := make([]byte, 1024)

	for {
		if count, err = reader.Read(part); err != nil {
			break
		}
		r.buffer.Write(part[:count])
	}
	if err != io.EOF {
		slog.Error("error reading file", "filePath", filePath, "error", err)
		return err
	}

	return nil
}
