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

package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"strings"
)

type ResourceConfig struct {
	Type         string
	NestedConfig map[string]interface{} `mapstructure:",remain"`
}

type Config struct {
	Source      ResourceConfig
	Destination ResourceConfig
	LogLevel    string
}

func (c *Config) GetLogLevel() slog.Level {
	var level slog.Level

	switch strings.ToLower(c.LogLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	return level
}

func (c *Config) DebugEnabled() bool {
	return strings.ToLower(c.LogLevel) == "debug"
}

func New() (*Config, error) {
	var c Config

	err := viper.Unmarshal(&c)
	if err != nil {
		slog.Error("unable to decode into struct", "error", err)
		return nil, err
	}

	return &c, nil
}
