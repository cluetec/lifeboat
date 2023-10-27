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
	"github.com/spf13/viper"
	"log/slog"
)

const Type = "filesystem"

type Config struct {
	path string
}

func New() (*Config, error) {
	var filesystemConfig Config

	err := viper.UnmarshalKey("destination", &filesystemConfig)

	if err != nil {
		slog.Error("unable to decode config into filesystem source config", slog.Any("error", err))
		return nil, err
	}

	return &filesystemConfig, nil
}
