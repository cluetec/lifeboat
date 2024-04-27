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
	"log/slog"
	"os"

	"github.com/cluetec/lifeboat/internal/config/validator"
)

type Reader struct {
	file *os.File
}

func NewReader(c *Config) (*Reader, error) {
	if err := validator.Validator.Struct(c); err != nil {
		return nil, err
	}

	slog.Debug("source config validated", "sourceType", Type, "config", c)

	f, err := os.Open(c.Path)
	if err != nil {
		return nil, err
	}

	return &Reader{file: f}, nil
}

func (r *Reader) Read(b []byte) (int, error) {
	slog.Debug("read got called", "sourceType", Type)
	return r.file.Read(b)
}

func (r *Reader) Close() error {
	slog.Debug("closing reader", "sourceType", Type)

	if r.file != nil {
		if err := r.file.Close(); err != nil {
			return err
		}
		r.file = nil
	}
	return nil
}
