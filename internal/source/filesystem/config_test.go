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

package filesystem

import (
	"testing"

	"github.com/cluetec/lifeboat/internal/config/validator"
	playValidator "github.com/go-playground/validator/v10"
)

func TestConfig_validation(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
		want    []validator.ExpectedFieldError
	}{
		{
			name: "Missing file or dir",
			config: Config{
				Path: "/source.txt",
			},
			wantErr: true,
			want: []validator.ExpectedFieldError{
				{Namespace: "Config.Path", Tag: "file|dir"},
			},
		},
		{
			name:    "Missing path",
			config:  Config{},
			wantErr: true,
			want: []validator.ExpectedFieldError{
				{Namespace: "Config.Path", Tag: "required"},
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
