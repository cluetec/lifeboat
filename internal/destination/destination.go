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

package destination

import (
	"errors"
	"io"
	"log/slog"

	"github.com/cluetec/lifeboat/internal/config"
	"github.com/cluetec/lifeboat/internal/destination/filesystem"
)

type Destination struct {
	Writer io.WriteCloser
}

func New(c *config.DestinationConfig) (*Destination, error) {
	d := Destination{}
	var err error

	switch {
	case c.Type == filesystem.Type:
		d.Writer, err = filesystem.NewWriter(&c.Filesystem)
	}
	if err != nil {
		slog.Error("error while initializing writer interface for destination system", "destinationType", c.Type, "error", err)
		return nil, err
	}

	if d.Writer == nil {
		return nil, errors.New("destination type not known")
	}

	return &d, nil
}
