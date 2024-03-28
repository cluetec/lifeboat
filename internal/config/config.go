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

type ResourceConfig map[string]any

type SourceConfig struct {
	Type           string
	ResourceConfig ResourceConfig `mapstructure:",remain"`
}

type DestinationConfig struct {
	Type           string
	ResourceConfig ResourceConfig `mapstructure:",remain"`
}

type Config struct {
	Source      SourceConfig
	Destination DestinationConfig
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

func New(cfgFilePath string) (*Config, error) {
	if cfgFilePath == "" {
		viper.AddConfigPath(".")
		viper.SetConfigFile("./config.yaml")
	} else {
		viper.SetConfigFile(cfgFilePath)
	}

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("error while reading in the configs: %w", err)
		return nil, err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		slog.Error("unable to decode into struct", "error", err)
		return nil, err
	}

	return &c, nil
}
