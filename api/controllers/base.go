package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/rs/zerolog/log"
)

type Response[D any, E any] struct {
	Data  D `json:"data"`
	Error E `json:"error"`
} //@Name Response

type ErrorObj struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
} //@Name Error

type Empty struct{} //@Name Empty

type BaseController struct {
	beego.Controller
}

func (b *BaseController) setResponseError(code int, message string) {
	response := Response[interface{}, ErrorObj]{
		Error: ErrorObj{Code: code, Message: message},
	}

	bts, _ := json.Marshal(&response)
	b.Controller.Ctx.Output.Header("Content-Type", context.ApplicationJSON)
	b.Controller.Ctx.Output.SetStatus(code)
	err := b.Controller.Ctx.Output.Body(bts)
	if err != nil {
		log.Err(err)
	}
}
