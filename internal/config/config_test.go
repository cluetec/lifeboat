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

package config

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Test default values", func(t *testing.T) {
		// given
		// create empty config file
		f := provideTempConfigFile(t, []byte(`
source:
  type: filesystem
  filesystem:
    path: /source/path.txt
destination:
  type: filesystem
  filesystem:
    path: /destination/path.txt`))
		defer removeTempConfigFile(f)

		want := &Config{
			LogLevel: "",
			Source: SourceConfig{
				Type: "filesystem",
				ResourceConfig: ResourceConfig{
					"filesystem": map[string]interface{}{
						"path": "/source/path.txt",
					},
				},
			},
			Destination: DestinationConfig{
				Type: "filesystem",
				ResourceConfig: ResourceConfig{
					"filesystem": map[string]interface{}{
						"path": "/destination/path.txt",
					},
				},
			},
		}

		// when
		got, err := New(f.Name())

		// then
		if err != nil {
			t.Fatalf("New() returned an error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("New() got = %v, want %v", got, want)
		}

		if got.GetLogLevel() != slog.LevelInfo {
			t.Errorf("New().GetLogLevel() got = %v, wanted = %v", got.GetLogLevel(), slog.LevelInfo)
		}
	})

	t.Run("Load config file", func(t *testing.T) {
		// given

		// write data to the temporary file
		var yamlExample = []byte(`
loglevel: warn
source:
  type: filesystem
destination:
  type: filesystem
`)
		f := provideTempConfigFile(t, yamlExample)
		defer removeTempConfigFile(f)

		want := &Config{
			LogLevel: "warn",
			Source: SourceConfig{
				Type:           "filesystem",
				ResourceConfig: nil,
			},
			Destination: DestinationConfig{
				Type:           "filesystem",
				ResourceConfig: nil,
			},
		}

		// when
		got, err := New(f.Name())

		// then
		if err != nil {
			t.Fatalf("New() returned an error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("New() got = %v, want %v", got, want)
		}
	})

	// ⚠️ If this tests fails, try to run it again with the build tag `viper_bind_struct` enabled.
	//	t.Run("Consideration of env variables", func(t *testing.T) {
	//		// given
	//		// create empty config file
	//		f := provideTempConfigFile(t, []byte(`
	//source:
	//  type: filesystem
	//destination:
	//  type: filesystem`))
	//		defer removeTempConfigFile(f)
	//
	//		// set environment variable
	//		defer os.Clearenv()
	//		_ = os.Setenv("LOGLEVEL", "warn")
	//		_ = os.Setenv("SOURCE_FILESYSTEM_PATH", "/source/path/from/env.txt")
	//		_ = os.Setenv("DESTINATION_FILESYSTEM_PATH", "/destination/path/from/env.txt")
	//
	//		want := &Config{
	//			LogLevel: "warn",
	//			Source: SourceConfig{
	//				Type: "filesystem",
	//				ResourceConfig: ResourceConfig{
	//					"filesystem": map[string]interface{}{
	//						"path": "/source/path/from/env.txt",
	//					},
	//				},
	//			},
	//			Destination: DestinationConfig{
	//				Type: "filesystem",
	//				ResourceConfig: ResourceConfig{
	//					"filesystem": map[string]interface{}{
	//						"path": "/destination/path/from/env.txt",
	//					},
	//				},
	//			},
	//		}
	//
	//		// when
	//		got, err := New(f.Name())
	//
	//		// then
	//		if err != nil {
	//			t.Fatalf("New() returned an error: %v", err)
	//		}
	//		if !reflect.DeepEqual(got, want) {
	//			t.Errorf("New() got = %v, want %v", got, want)
	//		}
	//	})

	t.Run("Consideration of env variables over config file", func(t *testing.T) {
		// given
		// write data to the temporary file
		var yamlExample = []byte(`
loglevel: warn
source:
  type: filesystem
  filesystem:
    path: /source/path.txt
destination:
  type: filesystem
  filesystem:
    path: /destination/path.txt
`)
		f := provideTempConfigFile(t, yamlExample)
		defer removeTempConfigFile(f)

		defer os.Clearenv()
		_ = os.Setenv("LOGLEVEL", "error")
		_ = os.Setenv("SOURCE_FILESYSTEM_PATH", "/source/path/from/env.txt")
		_ = os.Setenv("DESTINATION_FILESYSTEM_PATH", "/destination/path/from/env.txt")

		want := &Config{
			LogLevel: "error",
			Source: SourceConfig{
				Type: "filesystem",
				ResourceConfig: ResourceConfig{
					"filesystem": map[string]interface{}{
						"path": "/source/path/from/env.txt",
					},
				},
			},
			Destination: DestinationConfig{
				Type: "filesystem",
				ResourceConfig: ResourceConfig{
					"filesystem": map[string]interface{}{
						"path": "/destination/path/from/env.txt",
					},
				},
			},
		}

		// when
		got, err := New(f.Name())

		// then
		if err != nil {
			t.Fatalf("New() returned an error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("New() got = %v, want %v", got, want)
		}
	})
}

func TestConfig_GetLogLevel(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   slog.Level
	}{
		{name: "debug", config: Config{LogLevel: "debug"}, want: slog.LevelDebug},
		{name: "info", config: Config{LogLevel: "info"}, want: slog.LevelInfo},
		{name: "warn", config: Config{LogLevel: "warn"}, want: slog.LevelWarn},
		{name: "error", config: Config{LogLevel: "error"}, want: slog.LevelError},
		{name: "DEBUG", config: Config{LogLevel: "DEBUG"}, want: slog.LevelDebug},
		{name: "INFO", config: Config{LogLevel: "INFO"}, want: slog.LevelInfo},
		{name: "WARN", config: Config{LogLevel: "WARN"}, want: slog.LevelWarn},
		{name: "ERROR", config: Config{LogLevel: "ERROR"}, want: slog.LevelError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.config.GetLogLevel(); got != tt.want {
				t.Errorf("GetLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// provideTempConfigFile creates a temporary config file and writes the byte array to the file.
//
// Attention: Everywhere where this method is called, it's also necessary to execute removeTempConfigFile to close and
// remove the temp file again, which could be done e.g. in a defer block.
func provideTempConfigFile(t *testing.T, yamlPayload []byte) *os.File {
	// given
	// create and open a temporary file
	f, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("error while creating temp config file: %v", err)
	}

	if _, err := f.Write(yamlPayload); err != nil {
		t.Fatalf("error while writing to temp config file: %v", err)
	}
	if err := f.Close(); err != nil {
		t.Fatalf("error while closing file: %v", err)
	}

	return f
}

// removeTempConfigFile closes and removes the given file.
func removeTempConfigFile(f *os.File) {
	_ = f.Close()
	_ = os.Remove(f.Name())
}
