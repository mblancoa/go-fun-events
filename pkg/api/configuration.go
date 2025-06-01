package api

import (
	"github.com/rs/zerolog/log"
)

func SetupApiConfiguration() {
	log.Info().Msg("Initializing API configuration")
	initRouters()
}
