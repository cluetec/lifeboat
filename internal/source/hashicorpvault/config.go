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
	"log/slog"

	vault "github.com/hashicorp/vault/api"
)

const Type = "hashicorpvault"

type KubernetesAuth struct {
	RoleName string `validate:"required"`
}

type Config struct {
	Address        string         `validate:"required,http_url"`
	Token          string         `validate:"required_if=AuthMethod token"`
	AuthMethod     string         `validate:"required,oneof=token kubernetes"`
	KubernetesAuth KubernetesAuth `validate:"required_if=AuthMethod kubernetes,omitempty"`
}

// LogValue customizes how the `config` struct will be printed in the logs.
func (c *Config) LogValue() slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("address", c.Address))
	groupValues = append(groupValues, slog.String("authMethod", c.AuthMethod))

	if c.AuthMethod == "kubernetes" {
		groupValues = append(groupValues, slog.String("roleName", c.KubernetesAuth.RoleName))
	}

	return slog.GroupValue(groupValues...)
}

func (c *Config) GetHashiCorpVaultConfig() *vault.Config {
	config := vault.Config{
		Address: c.Address,
	}

	return &config
}
