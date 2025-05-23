package api

import (
	"github.com/astaxie/beego"
	_ "github.com/mblancoa/go-fun-events/docs"
	"github.com/rs/zerolog/log"
	swagger "github.com/weblfe/beego-swagger"
)

func initRouters() {
	log.Info().Msg("Initializing events api routes")
	beego.Get("/swagger/*", swagger.Handler)

	w := WebContext
	beego.Router("/search", w.EventController, "get:Search")
}
