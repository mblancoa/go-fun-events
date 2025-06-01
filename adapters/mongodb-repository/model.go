package mongodb_repository

import (
	"github.com/devfeel/mapper"
	"github.com/google/uuid"
	"github.com/mblancoa/go-fun-events/pkg/core"
	"time"
)

const (
	EventsCollection string = "events"
)

type EventDB struct {
	Id       uuid.UUID `bson:"_id"`
	ProvId   string    `bson:"prov_id"`
	Title    string    `bson:"title"`
	StartsAt time.Time `bson:"starts_at"`
	EndsAt   time.Time `bson:"ends_at"`
	MaxPrice float64   `bson:"max_price"`
	MinPrice float64   `bson:"min_price"`
}

func MapToEvent(db *EventDB) (*core.Event, error) {
	event := &core.Event{}
	err := mapper.Mapper(db, event)
	if err != nil {
		return &core.Event{}, err
	}
	return event, nil
}

func MapToEventArray(dbArray []*EventDB) ([]*core.Event, error) {
	var events []*core.Event
	for _, db := range dbArray {
		event, err := MapToEvent(db)
		if err != nil {
			return []*core.Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}

func MapToEventDB(event *core.Event) (*EventDB, error) {
	db := &EventDB{}
	err := mapper.Mapper(event, db)
	if err != nil {
		return &EventDB{}, err
	}
	return db, nil
}

func MapToEventDBArray(events []*core.Event) ([]*EventDB, error) {
	var dbArray []*EventDB
	for _, event := range events {
		db, err := MapToEventDB(event)
		if err != nil {
			return []*EventDB{}, err
		}
		dbArray = append(dbArray, db)
	}
	return dbArray, nil
}
