package ports

import (
	"github.com/mblanco/Go-Acme-events/core/domain"
)

var ProviderContext *providerContext = &providerContext{}

type EventProvider interface {
	GetEvents() ([]*domain.Event, error)
}

type providerContext struct {
	EventProvider EventProvider
}
