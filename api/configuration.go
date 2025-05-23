package api

import (
	"github.com/mblancoa/go-fun-events/api/controllers"
	"github.com/mblancoa/go-fun-events/core"
	"github.com/rs/zerolog/log"
)

var WebContext *webContext = &webContext{}

type webContext struct {
	EventController *controllers.EventController
}

func SetupApiConfiguration() {
	log.Info().Msg("Initializing API configuration")
	setupWebContext()
	initRouters()
}

func setupWebContext() {
	log.Info().Msg("Creating the web context")
	c := core.DomainContext
	WebContext.EventController = controllers.NewEventController(c.EventService)
}
