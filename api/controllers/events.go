package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/validation"
	"github.com/mblancoa/go-fun-events/core"
	"github.com/mblancoa/go-fun-events/tools"
	"net/http"
	"time"
)

const (
	StartsAtQueryParam = "starts_at"
	EndsAtQueryParam   = "ends_at"
)

type EventController struct {
	BaseController
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
	startsAt, endsAt, err := getAndValidateSearchQueryParams(e.Ctx)
	if err != nil {
		e.setResponseError(http.StatusBadRequest, err.Error())
		return
	}

	events, err := core.DomainContext.EventService.GetEvents(startsAt, endsAt)
	if err != nil {
		e.setResponseError(http.StatusInternalServerError, err.Error())
		return
	}

	eventList := mapEventsToEventList(tools.FromPointerArray(events))
	response := Response[EventList, interface{}]{Data: eventList}

	e.setResponseBody(http.StatusOK, response)
}

func getAndValidateSearchQueryParams(ctx *context.Context) (time.Time, time.Time, error) {
	valid := validation.Validation{}
	var startsAt, endsAt time.Time

	start := ctx.Input.Query(StartsAtQueryParam)
	valid.Required(start, StartsAtQueryParam)
	end := ctx.Input.Query(EndsAtQueryParam)
	valid.Required(end, EndsAtQueryParam)
	if valid.HasErrors() {
		return time.Time{}, time.Time{}, validationError("Bad request", valid.Errors)
	}

	startsAt, err := time.Parse(DateTimeLayout, start)
	if err != nil {
		_ = valid.SetError(StartsAtQueryParam, err.Error())
	}
	endsAt, err = time.Parse(DateTimeLayout, end)
	if err != nil {
		_ = valid.SetError(EndsAtQueryParam, err.Error())
	}
	if valid.HasErrors() {
		return time.Time{}, time.Time{}, validationError("Bad request", valid.Errors)
	}

	return startsAt, endsAt, nil
}
