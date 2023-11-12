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
	"bytes"
	"compress/gzip"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	globalConfig "github.com/cluetec/lifeboat/internal/config"
)

type Reader struct {
	file        *os.File
	archivedDir *bytes.Buffer
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

	r := &Reader{}

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

	if r.file != nil {
		return r.file.Read(b)
	} else {
		return r.archivedDir.Read(b)
	}
}

func (r *Reader) Close() error {
	slog.Debug("closing filesystem reader")

	if r.file != nil {
		if err := r.file.Close(); err != nil {
			return err
		}
		r.file = nil
	}
	return nil
}

// prepareDir reads in the source directory for the backup as a tar archive and sets
// the contents as a bytes.Buffer to the Reader so it can be written during the backup
func (r *Reader) prepareDir(srcPath string) error {
	slog.Debug("preparing filesystem directory backup")
	var buf bytes.Buffer

	mw := io.MultiWriter(&buf)

	gw := gzip.NewWriter(mw)
	tw := tar.NewWriter(gw)

	err := filepath.WalkDir(srcPath, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// exclude root dir
		if path == srcPath {
			return nil
		}

		fileInfo, err := entry.Info()
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(fileInfo, entry.Name())
		if err != nil {
			return err
		}

		// preserve folder structure inside tar archive, f.e. for files in nested directories
		header.Name = strings.TrimPrefix(strings.Replace(path, srcPath, "", -1), string(filepath.Separator))

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !entry.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	if err := tw.Close(); err != nil {
		return err
	}

	if err := gw.Close(); err != nil {
		return err
	}

	r.archivedDir = &buf

	return nil
}

// prepareFile opens the source file with the given path and sets it at the reader
// in preparation for the backup
func (r *Reader) prepareFile(srcPath string) error {
	slog.Debug("preparing filesystem file backup")
	f, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	r.file = f

	return nil
}
