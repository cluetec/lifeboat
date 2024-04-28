/*
 * Copyright 2024 cluetec GmbH
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
	"reflect"
	"testing"

	"github.com/cluetec/lifeboat/internal/config/validator"
	playValidator "github.com/go-playground/validator/v10"
	vault "github.com/hashicorp/vault/api"
)

func TestConfig_validation(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
		want    []validator.ExpectedFieldError
	}{
		{
			name: "Successful validation",
			config: Config{
				Address: "http://localhost:8200",
				Token:   "root",
			},
			wantErr: false,
			want:    []validator.ExpectedFieldError{},
		},
		{
			name:    "Missing address & token",
			config:  Config{},
			wantErr: true,
			want: []validator.ExpectedFieldError{
				{Namespace: "Config.Address", Tag: "required"},
				{Namespace: "Config.Token", Tag: "required"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validator.Struct(tt.config)

			if (err != nil && !tt.wantErr) ||
				(err == nil && tt.wantErr) {
				t.Errorf("Config{} error = %v, wantErr %v", err, tt.wantErr)
			}

			// Iterate over validation errors
			// The if statement is necessary, because else we run into an panic
			if err != nil && tt.wantErr {
				for _, e := range err.(playValidator.ValidationErrors) {
					validError := false

					// Iterate over expected field errors to identify if error was expected
					for _, expectedFieldError := range tt.want {
						if validator.ValidateFieldError(e, expectedFieldError) {
							validError = true
							break
						}
					}

					if !validError {
						t.Errorf("Config{} got = %v, want %v", e, tt.want)
					}
				}
			}
		})
	}
}

func TestConfig_GetHashiCorpVaultConfig(t *testing.T) {
	// given
	c := &Config{
		Address: "http://localhost:8200",
	}
	want := &vault.Config{Address: c.Address}

	// when
	got := c.GetHashiCorpVaultConfig()

	// then
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetHashiCorpVaultConfig() = %v, want %v", got, want)
	}
}

func TestConfig_LogValue(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   slog.Value
	}{
		{
			name: "Successful log value",
			config: Config{
				Address: "http://localhost:8200",
				Token:   "root",
			},
			want: slog.GroupValue(
				slog.Attr{Key: "address", Value: slog.StringValue("http://localhost:8200")},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.LogValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
