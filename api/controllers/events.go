package controllers

import (
	"github.com/mblancoa/go-fun-events/core"
	"net/http"
)

type EventController struct {
	BaseController
	eventService core.EventService
}

func NewEventController(eventService core.EventService) *EventController {
	return &EventController{eventService: eventService}
}

// @Title Events searching
// @Description Search events between two dates
// @Accept       mpfd
// @Produce      json
// @Param        starts_at   query      string  true  "2018-05-29T21:00:00"
// @Param        ends_at     query      string  true  "2025-06-30T21:00:00"
// @Success 200 {object} controllers.Request[controllers.EventList]
// @Failure      400  {object}  controllers.Request[interface{}]
// @Failure      500  {object}  controllers.Request[interface{}]
// @Router /search [get]

func (e *EventController) Search() {
	//TODO implement me
	e.setResponseError(http.StatusConflict, "probando error")
}
