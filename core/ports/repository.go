package ports

import (
	"github.com/mblanco/Go-Acme-events/core/domain"
	"time"
)

var RepositoryContext *repositoryContext = &repositoryContext{}

type EventRepository interface {
	FindByStartAfterAndEndBefore(from, to time.Time) ([]*domain.Event, error)
	Update(toUpdate []*domain.Event) error
	InsertOrUpdate(toInsert []*domain.Event) error
}

type repositoryContext struct {
	EventRepository EventRepository
}
