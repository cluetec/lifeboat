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
	"errors"
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
	var err error

	switch {
	case c.Type == filesystem.Type:
		s.Reader, err = filesystem.NewReader(&c.ResourceConfig)
	}
	if err != nil {
		slog.Error("error while initializing reader interface for source system", "sourceType", c.Type, "error", err)
		return nil, err
	}

	if s.Reader == nil {
		return nil, errors.New("source type not known")
	}

	return &s, nil
}

//
//case c.Type == hashicorpvault.Type:
//s.Reader, err = hashicorpvault.NewReader(&c.ResourceConfig)
//if err != nil {
//slog.Error("error while initializing reader interface for source system", "sourceType", hashicorpvault.Type, "error", err)
//return nil, err
//}
//return &s, nil
