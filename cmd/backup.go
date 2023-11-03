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
	"log/slog"
	"os"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Execute the backup.",
	Long:  "Execute the backup. Used config can be overridden by providing arguments.",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.New()
		if err != nil {
			slog.Error("error while initializing config", "error", err)
			os.Exit(1)
		}

		logging.InitSlog(c.GetLogLevel())

		slog.Debug("global config loaded", slog.Any("globalConfig", c))
		slog.Debug("log level set", slog.String("logLevel", c.LogLevel))

		slog.Debug("start of backup command")

		source.Prepare(c.Source)
		destination.Prepare(c.Destination)

		slog.Info("TODO: Do backup")

		slog.Debug("end of backup command")
	},
}
