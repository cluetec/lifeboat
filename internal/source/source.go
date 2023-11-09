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

package source

import (
	"github.com/cluetec/lifeboat/internal/config"
	"github.com/cluetec/lifeboat/internal/source/filesystem"
	"io"
	"log/slog"
)

type Source struct {
	Reader io.ReadCloser
}

func New(c config.SourceConfig) (*Source, error) {
	s := Source{}
	if c.Type == filesystem.Type {
		filesystemConfig, err := filesystem.NewConfig(c.ResourceConfig)
		if err != nil {
			slog.Error("error while initializing filesystem source config", "error", err)
			return nil, err
		}

		slog.Debug("filesystem source config loaded", "config", filesystemConfig)

		s.Reader, err = filesystem.NewReader(filesystemConfig)
		if err != nil {
			slog.Error("error while initializing reader interface for filesystem source", "error", err)
			return nil, err
		}
	}

	return &s, nil
}
