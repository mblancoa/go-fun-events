package core

import "github.com/rs/zerolog/log"

type SupplyService interface {
	FetchEventsFromProvider()
}

type supplyService struct {
	eventService  EventService
	eventProvider EventProvider
}

func NewSupplyService(eventService EventService, eventProvider EventProvider) SupplyService {
	return &supplyService{
		eventService:  eventService,
		eventProvider: eventProvider,
	}
}

func (s *supplyService) FetchEventsFromProvider() {
	log.Info().Msgf("Fetching events from provider...")
	events, err := s.eventProvider.GetEvents()
	if err != nil {
		log.Debug().Msgf("Exception fetching events: %s", err.Error())
		return
	}
	log.Info().Msgf("%v events have been fetched", len(events))

	s.eventService.UpdateEvents(events)
}
