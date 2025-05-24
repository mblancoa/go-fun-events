package api

import (
	"github.com/astaxie/beego"
	"github.com/mblancoa/go-fun-events/api/controllers"
	_ "github.com/mblancoa/go-fun-events/docs"
	"github.com/rs/zerolog/log"
	swagger "github.com/weblfe/beego-swagger"
)

func initRouters() {
	log.Info().Msg("Initializing events api routes")
	beego.Get("/swagger/*", swagger.Handler)

	beego.Router("/search", &controllers.EventController{}, "get:Search")
}
