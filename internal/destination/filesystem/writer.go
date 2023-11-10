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
	"errors"
	globalConfig "github.com/cluetec/lifeboat/internal/config"
	"log/slog"
	"os"
)

type Writer struct {
	file *os.File
}

func NewWriter(c *globalConfig.ResourceConfig) (*Writer, error) {
	rc, err := newConfig(c)
	if err != nil {
		slog.Error("error while initializing filesystem destination config", "error", err)
		return nil, err
	}

	slog.Debug("filesystem destination config loaded", "config", rc)

	// Check if destination file already exists
	_, err = os.Stat(rc.Path)
	if err == nil {
		return nil, errors.New("destination file already exists")
	} else if !errors.Is(err, os.ErrNotExist) {
		slog.Error("error while checking if destination file already exists", "error", err)
		return nil, err
	}

	// Create file
	f, err := os.Create(rc.Path)
	if err != nil {
		return nil, err
	}

	return &Writer{file: f}, nil
}

func (w *Writer) Write(b []byte) (int, error) {
	slog.Debug("filesystem destination write got called")
	return w.file.Write(b)
}

func (w *Writer) Close() error {
	slog.Debug("closing filesystem writer")

	if w.file != nil {
		if err := w.file.Close(); err != nil {
			return err
		}
		w.file = nil
	}
	return nil
}
