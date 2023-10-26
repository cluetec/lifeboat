/*
 * Copyright (c) 2023 ZF Friedrichshafen AG
 */

package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"testing"
)

func TestConfig_GetLogLevel(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   zerolog.Level
	}{
		{name: "trace", config: Config{LogLevel: "trace"}, want: zerolog.TraceLevel},
		{name: "debug", config: Config{LogLevel: "debug"}, want: zerolog.DebugLevel},
		{name: "info", config: Config{LogLevel: "info"}, want: zerolog.InfoLevel},
		{name: "warn", config: Config{LogLevel: "warn"}, want: zerolog.WarnLevel},
		{name: "error", config: Config{LogLevel: "error"}, want: zerolog.ErrorLevel},
		{name: "TRACE", config: Config{LogLevel: "TRACE"}, want: zerolog.TraceLevel},
		{name: "DEBUG", config: Config{LogLevel: "DEBUG"}, want: zerolog.DebugLevel},
		{name: "INFO", config: Config{LogLevel: "INFO"}, want: zerolog.InfoLevel},
		{name: "WARN", config: Config{LogLevel: "WARN"}, want: zerolog.WarnLevel},
		{name: "ERROR", config: Config{LogLevel: "ERROR"}, want: zerolog.ErrorLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate := validator.New()
			if err := validate.Struct(tt.config); err != nil {
				t.Errorf("Config not valid: %v", err)
			}

			if got := tt.config.GetLogLevel(); got != tt.want {
				t.Errorf("GetLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
