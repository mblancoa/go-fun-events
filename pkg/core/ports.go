package core

import "time"

// ==================================
// Provider port definition
// ==================================

var ProviderContext *providerContext = &providerContext{}

type EventProvider interface {
	GetEvents() ([]*Event, error)
}

type providerContext struct {
	EventProvider EventProvider
}

// ==================================
// Repository port definition
// ==================================

var RepositoryContext *repositoryContext = &repositoryContext{}

type EventRepository interface {
	FindByStartAfterAndEndBefore(from, to time.Time) ([]*Event, error)
	Update(toUpdate []*Event) error
	InsertOrUpdate(toInsert []*Event) error
}

type repositoryContext struct {
	EventRepository EventRepository
}
