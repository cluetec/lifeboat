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
	"log/slog"
	"strings"

	"github.com/cluetec/lifeboat/internal/config/validator"
	destFilesystem "github.com/cluetec/lifeboat/internal/destination/filesystem"
	srcFilesystem "github.com/cluetec/lifeboat/internal/source/filesystem"
	"github.com/cluetec/lifeboat/internal/source/hashicorpvault"
	"github.com/spf13/viper"
)

type ResourceConfig map[string]any

type SourceConfig struct {
	Type           string                `validate:"required,oneof=filesystem hashicorpvault"`
	Filesystem     srcFilesystem.Config  `validate:"omitempty"`
	HashiCorpVault hashicorpvault.Config `validate:"omitempty"`
}

type DestinationConfig struct {
	Type       string                `validate:"required,oneof=filesystem"`
	Filesystem destFilesystem.Config `validate:"omitempty"`
}

type Config struct {
	Source      SourceConfig
	Destination DestinationConfig
	LogLevel    string `validate:"omitempty,oneof=debug DEBUG info INFO warn WARN error ERROR"`
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
		viper.SetConfigFile("config.yaml")
	} else {
		viper.SetConfigFile(cfgFilePath)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		slog.Error("unable to decode into struct", "error", err)
		return nil, err
	}

	if err := validator.Validator.Struct(c); err != nil {
		return nil, err
	}

	return &c, nil
}
