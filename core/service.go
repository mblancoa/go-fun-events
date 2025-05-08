package core

import (
	"github.com/mblanco/Go-Acme-events/core/domain"
	"github.com/mblanco/Go-Acme-events/core/ports"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

type EventService interface {
	GetEvents(startsAt, endsAt time.Time) ([]*domain.Event, error)
}

func NewEventService(
	timeToFeed time.Duration,
	eventRepository ports.EventRepository,
	eventProvider ports.EventProvider,
) EventService {
	eventService := eventService{
		timeToFeed:      timeToFeed,
		eventRepository: eventRepository,
		eventProvider:   eventProvider,
	}
	return &eventService
}

type eventService struct {
	lastUpdating time.Time
	locker       *sync.Mutex

	timeToFeed      time.Duration
	eventRepository ports.EventRepository
	eventProvider   ports.EventProvider
}

func (e *eventService) GetEvents(startsAt, endsAt time.Time) ([]*domain.Event, error) {
	e.updateEventsSynchronously()
	return e.eventRepository.FindByStartAfterAndEndBefore(startsAt, endsAt)
}

func (e *eventService) updateEventsSynchronously() {
	if e.locker.TryLock() {
		defer e.locker.Unlock()
		e.updateEvents()
	}
}
func (e *eventService) updateEvents() {
	now := time.Now()
	if e.lastUpdating.IsZero() || e.lastUpdating.Add(e.timeToFeed).After(now) {
		events, err := e.eventProvider.GetEvents()
		if err != nil {
			log.Debug().Msgf("Exception fetching events: %s", err.Error())
			return
		}
		if events != nil && len(events) != 0 {
			var onlineEvents []*domain.Event
			var restEvents []*domain.Event
			for _, event := range events {
				if event.IsOnlineSale {
					onlineEvents = append(onlineEvents, event)
				} else {
					restEvents = append(restEvents, event)
				}
			}
			err = e.eventRepository.InsertOrUpdate(onlineEvents)
			if err != nil {
				log.Warn().Msgf("Exception inserting and updating events: %s", err.Error())
				return
			}
			err = e.eventRepository.Update(restEvents)
			if err != nil {
				log.Warn().Msgf("Exception updating events: %s", err.Error())
				return
			}
			e.lastUpdating = time.Now()
		}
	}
}
