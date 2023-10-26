package sources

import (
	"github.com/rs/zerolog/log"
	"io"
)

func CloseResponseBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatal().
			Stack().
			Err(err).Msg("Error while closing response body")
	}
}
