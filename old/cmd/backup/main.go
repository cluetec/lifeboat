package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
	"zf.com/dos/backup/internal/config"
	"zf.com/dos/backup/internal/destination/azureblob"
	"zf.com/dos/backup/internal/logging"
	"zf.com/dos/backup/internal/sources"
	"zf.com/dos/backup/internal/sources/hashicorpvault"
)

func main() {
	initZerolog()

	c, err := config.NewConfig()
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("Failed loading config")
	}

	zerolog.SetGlobalLevel(c.GetLogLevel())

	log.Info().Interface("config", c).Msg("Config loaded")

	vaultClient, err := hashicorpvault.InitClient(&c.Source.HashiCorpVault)
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("Failed to initialize HashiCorp Vault client")
	}

	// Create and get the snapshot via REST call
	// Timout needs to be adjusted eventually
	response, err := hashicorpvault.Backup(vaultClient)
	logging.GetHttpTracingContext()
	log.With().Dict("request", logEvent).Logger()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).
			Msg("Create snapshot request failed")
	}
	defer sources.CloseResponseBody(response.Body)

	// Write snapshot to file
	// Create file
	filename := time.Now().Format("2006-01-02_1504") + "_vault.snap"

	logger := log.With().
		Str("filename", filename).
		Str("request.url", response.Request.URL.String()).
		Dict("response", zerolog.Dict().
			Int("statusCode", response.StatusCode)).
		Logger()

	//// Store backup to filesystem
	//fileDestination := filesystem.Destination{
	//	Logger: logger,
	//}
	//fileDestination.Store(response, filename)

	// Store backup to azure blob
	blobDestination := azureblob.Destination{
		Logger: logger,
		Config: c.Destination.AzureBlob,
	}
	blobDestination.Store(response, filename)
}
