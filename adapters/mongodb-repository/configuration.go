package mongodb_repository

import "github.com/rs/zerolog/log"

func SetupRepositoryConfiguration() {
	//Todo
	log.Info().Msg("Initializing mongodb repository configuration")
	setupRepositoryContext()
}
func setupRepositoryContext() {
	log.Info().Msg("Creating the mongodb repository context")
	//Todo
	//core.RepositoryContext.EventRepository=
}
