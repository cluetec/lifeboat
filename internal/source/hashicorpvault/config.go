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

package hashicorpvault

import (
	globalConfig "github.com/cluetec/lifeboat/internal/config"
	"github.com/go-playground/validator/v10"
	vault "github.com/hashicorp/vault/api"
	"github.com/mitchellh/mapstructure"
	"log/slog"
)

const Type = "hashicorpvault"

type config struct {
	Address string `validate:"http_url,required"`
	Token   string `validate:"required"`
}

var validate *validator.Validate

// newConfig provides the specific `config` struct. Therefor it takes the generic `globalConfig.ResourceConfig` and
// decodes it into the `config` struct and validates the values.
func newConfig(rc *globalConfig.ResourceConfig) (*config, error) {
	var c config

	err := mapstructure.Decode(rc, &c)
	if err != nil {
		slog.Error("unable to decode config into HashiCorp Vault source config", "error", err)
		return nil, err
	}

	validate = validator.New()
	if err := validate.Struct(c); err != nil {
		return nil, err
	}

	return &c, nil
}

// LogValue customizes how the `config` struct will be printed in the logs.
func (c *config) LogValue() slog.Value {
	return slog.GroupValue(slog.String("address", c.Address), slog.String("token", "***"))
}

// GetHashiCorpVaultConfig was implement in regard to the
// `vault.DefaultConfig()` method. While the implementation the
// `config.ReadEnvironment` was left of, to avoid the usage of additional
// environment variables like `VAULT_ADDRESS`.
func (c *config) GetHashiCorpVaultConfig() *vault.Config {
	config := vault.Config{
		Address: c.Address,
	}

	return &config
}
