package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/rs/zerolog/log"
)

type Response[T any] struct {
	Data  T             `json:"data"`
	Error ErrorResponse `json:"error"`
}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseController struct {
	beego.Controller
}

func (b *BaseController) setResponseError(code int, message string) {
	response := Response[interface{}]{
		Error: ErrorResponse{Code: code, Message: message},
	}

	bts, _ := json.Marshal(&response)
	b.Controller.Ctx.Output.Header("Content-Type", context.ApplicationJSON)
	b.Controller.Ctx.Output.SetStatus(code)
	err := b.Controller.Ctx.Output.Body(bts)
	if err != nil {
		log.Err(err)
	}
}
