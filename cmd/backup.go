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

package cmd

import (
	"github.com/cluetec/lifeboat/internal/config"
	"github.com/cluetec/lifeboat/internal/destination"
	"github.com/cluetec/lifeboat/internal/logging"
	"github.com/cluetec/lifeboat/internal/source"
	"github.com/spf13/cobra"
	"io"
	"log/slog"
)

var cfgFilePath string

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Execute the backup.",
	Long:  "Execute the backup. Used config can be overridden by providing arguments.",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := config.New(cfgFilePath)
		if err != nil {
			slog.Error("error while initializing config", "error", err)
			return err
		}

		logging.InitSlog(c.GetLogLevel())

		slog.Debug("global config loaded", slog.Any("globalConfig", c))
		slog.Debug("log level set", slog.String("logLevel", c.LogLevel))

		slog.Debug("start of backup command")

		s, err := source.New(c.Source)
		if err != nil {
			slog.Error("error while initializing source", "error", err)
			return err
		}
		defer func() {
			err := s.Reader.Close()
			if err != nil {
				slog.Error("error while closing source reader", "error", err)
			}
		}()

		d, err := destination.New(c.Destination)
		if err != nil {
			slog.Error("error while initializing destination", "error", err)
			return err
		}
		defer func() {
			err := d.Writer.Close()
			if err != nil {
				slog.Error("error while closing destination writer", "error", err)
			}
		}()

		n, err := io.Copy(d.Writer, s.Reader)
		if err != nil {
			slog.Error("error while doing the backup", "error", err)
			return err
		}

		slog.Info("Backup successful", "writtenBytes", n)
		return nil
	},
}

func init() {
	backupCmd.PersistentFlags().StringVarP(
		&cfgFilePath,
		"config",
		"c",
		"",
		"path to config file (default: ./config.yaml)",
	)
}
