package controllers

import (
	"github.com/google/uuid"
	"github.com/mblancoa/go-fun-events/core"
)

const (
	DateTimeLayout = "2006-01-02T15:04:05"
	DateLayout     = "2006-01-02"
	TimeLayout     = "15:04"
)

type EventDto struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	StartDate string    `json:"start_date" example:"2018-05-29"`
	StartTime string    `json:"start_time" example:"21:00"`
	EndDate   string    `json:"end_date" example:"2025-06-30"`
	EndTime   string    `json:"end_time" example:"21:00"`
	MinPrice  float64   `json:"min_price"`
	MaxPrice  float64   `json:"max_price"`
} //@Name EventSummary

type EventList struct {
	Events []EventDto `json:"events"`
} //@Name EventList

func mapEventsToEventList(events []core.Event) EventList {
	n := len(events)
	dto := make([]EventDto, n)
	for i := 0; i < n; i++ {
		dto[i] = mapEventToEventDto(events[i])
	}
	return EventList{Events: dto}
}

func mapEventToEventDto(event core.Event) EventDto {
	return EventDto{
		Id:        event.Id,
		Title:     event.Title,
		StartDate: event.StartsAt.Format(DateLayout),
		StartTime: event.StartsAt.Format(TimeLayout),
		EndDate:   event.EndsAt.Format(DateLayout),
		EndTime:   event.EndsAt.Format(TimeLayout),
		MinPrice:  event.MinPrice,
		MaxPrice:  event.MaxPrice,
	}
}
