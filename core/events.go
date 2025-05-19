package core

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
)

type Event struct {
	Id           uuid.UUID
	ProvId       string
	Title        string
	IsOnlineSale bool
	StartsAt     time.Time
	EndsAt       time.Time
	MaxPrice     float64
	MinPrice     float64
}

type EventService interface {
	GetEvents(startsAt, endsAt time.Time) ([]*Event, error)
	UpdateEvents(events []*Event)
}

func NewEventService(
	eventRepository EventRepository,
) EventService {
	eventService := eventService{
		eventRepository: eventRepository,
	}
	return &eventService
}

type eventService struct {
	eventRepository EventRepository
}

func (e *eventService) GetEvents(startsAt, endsAt time.Time) ([]*Event, error) {
	return e.eventRepository.FindByStartAfterAndEndBefore(startsAt, endsAt)
}

func (e *eventService) UpdateEvents(events []*Event) {
	if events != nil && len(events) != 0 {
		var onlineEvents []*Event
		var restEvents []*Event
		for _, event := range events {
			if event.IsOnlineSale {
				onlineEvents = append(onlineEvents, event)
			} else {
				restEvents = append(restEvents, event)
			}
		}
		err := e.eventRepository.InsertOrUpdate(onlineEvents)
		if err != nil {
			log.Warn().Msgf("Exception inserting and updating events: %s", err.Error())
			return
		}
		err = e.eventRepository.Update(restEvents)
		if err != nil {
			log.Warn().Msgf("Exception updating events: %s", err.Error())
			return
		}
	}
}
