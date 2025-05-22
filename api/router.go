package api

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/rs/zerolog/log"
)

func initRouter() {
	log.Info().Msg("Initializing events api routes")
	w := WebContext
	beego.Router("/search", w.EventController, "get:Search")
}
