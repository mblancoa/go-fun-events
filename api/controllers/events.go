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

// @Summary		Events searching
// @Tags		events
// @Title		Events searching
// @Description	Searches events between two dates
// @Accept		mpfd
// @Produce		json
// @Param		starts_at	query string true "2018-05-29T21:00:00"
// @Param		ends_at		query string true "2025-06-30T21:00:00"
// @Success		200	{object} Response[EventList,Empty]	"Success"
// @Failure		400	{object} Response[Empty, ErrorObj]		"Bad Request"
// @Failure		500	{object} Response[Empty, ErrorObj]		"Internal error"
// @Router		/search [get]
func (e *EventController) Search() {
	//TODO implement me
	e.setResponseError(http.StatusConflict, "probando error")
}
