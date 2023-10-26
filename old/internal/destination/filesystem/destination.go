package filesystem

import (
	vault "github.com/hashicorp/vault/api"
	"github.com/rs/zerolog"
	"io"
	"os"
)

type Destination struct {
	Logger zerolog.Logger
}

func (d *Destination) Store(response *vault.Response, filename string) {
	out, err := os.Create(filename)
	if err != nil {
		d.Logger.Fatal().
			Stack().
			Err(err).
			Msg("Could not create snapshot file on disc")
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			d.Logger.Fatal().
				Stack().
				Err(err).Msg("Error while closing new backup file")
		}
	}(out)

	writtenBytes, err := io.Copy(out, response.Body)
	logger := d.Logger.With().
		Int64("writtenBytes", writtenBytes).
		Logger()
	if err != nil {
		logger.Fatal().
			Stack().
			Err(err).
			Msg("Error while writing snapshot to file")
	}

	logger.Info().
		Msg("Successfully wrote to file")
}
