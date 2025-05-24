package controllers

import (
	"encoding/json"
	errors2 "errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/rs/zerolog/log"
	"net/http"
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

func (b *BaseController) setResponseBody(code int, response interface{}) {
	bts, _ := json.Marshal(&response)
	b.Controller.Ctx.Output.Header("Content-Type", context.ApplicationJSON)
	err := b.Controller.Ctx.Output.Body(bts)
	if err != nil {
		log.Err(err)
		b.setResponseError(http.StatusInternalServerError, err.Error())
	}
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

func validationError(message string, errors []*validation.Error) error {
	for _, e := range errors {
		message += fmt.Sprintf("\n[%s] property %s", e.Key, e.Error())
	}
	return errors2.New(message)
}
