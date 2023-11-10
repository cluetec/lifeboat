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
	globalConfig "github.com/cluetec/lifeboat/internal/config"
	"log/slog"
	"os"
)

type Reader struct {
	file *os.File
}

func NewReader(c *globalConfig.ResourceConfig) (*Reader, error) {
	rc, err := newConfig(c)
	if err != nil {
		slog.Error("error while initializing filesystem source config", "error", err)
		return nil, err
	}

	slog.Debug("filesystem source config loaded", "config", rc)

	f, err := os.Open(rc.Path)
	if err != nil {
		return nil, err
	}

	return &Reader{file: f}, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	slog.Debug("filesystem source read got called")
	return r.file.Read(b)
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
