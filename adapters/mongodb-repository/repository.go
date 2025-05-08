package mongodb_repository

import (
	"github.com/mblanco/Go-Acme-events/core/domain"
	"github.com/mblanco/Go-Acme-events/core/ports"
	"time"
)

type mongodbEventRepository struct {
}

func NewEventRepository() ports.EventRepository {
	return &mongodbEventRepository{}
}

func (m *mongodbEventRepository) FindByStartAfterAndEndBefore(from, to time.Time) ([]*domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mongodbEventRepository) Update(toUpdate []*domain.Event) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongodbEventRepository) InsertOrUpdate(toInsert []*domain.Event) error {
	//TODO implement me
	panic("implement me")
}
