package api

import (
	"github.com/astaxie/beego"
	_ "github.com/mblancoa/go-fun-events/docs"
	"github.com/mblancoa/go-fun-events/pkg/api/controllers"
	"github.com/rs/zerolog/log"
	swagger "github.com/weblfe/beego-swagger"
)

func initRouters() {
	log.Info().Msg("Initializing events api routes")
	beego.Get("/swagger/*", swagger.Handler)

	beego.Router("/search", &controllers.EventController{}, "get:Search")
}
