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
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"log/slog"
)

const Type = "filesystem"

type metaConfig struct {
	Filesystem config
}

type config struct {
	Path string `validate:"filepath,required"`
}

var validate *validator.Validate

func newConfig(rc *globalConfig.ResourceConfig) (*config, error) {
	var c metaConfig

	err := mapstructure.Decode(rc, &c)

	if err != nil {
		slog.Error("unable to decode config into filesystem destination config", "error", err)
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(c); err != nil {
		return nil, err
	}

	return &c.Filesystem, nil
}
