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
	"github.com/cluetec/lifeboat/internal/config"
	"github.com/cluetec/lifeboat/internal/destination/filesystem"
	"log/slog"
	"os"
)

func Prepare(c config.DestinationConfig) {
	if c.Type == filesystem.Type {
		filesystemConfig, err := filesystem.New(c.ResourceConfig)
		if err != nil {
			slog.Error("error while initializing filesystem destination config", "error", err)
			os.Exit(1)
		}

		slog.Debug("filesystem destination config loaded", "config", filesystemConfig)
	}
}