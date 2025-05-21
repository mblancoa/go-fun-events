package mongodb_repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mblancoa/go-fun-events/core"
	"github.com/mblancoa/go-fun-events/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//go:generate repogen -dest=mongodbeventrepository_impl.go -model=EventDB -repo=MongoDbEventRepository
type MongoDbEventRepository interface {
	InsertOne(ctx context.Context, event *EventDB) (interface{}, error)
	FindByProvId(ctx context.Context, provId string) (*EventDB, error)
	FindByStartsAtGreaterThanEqualAndEndsAtLessThanEqual(ctx context.Context, from, to time.Time) ([]*EventDB, error)
}

type eventRepository struct {
	collection *mongo.Collection
	delegate   MongoDbEventRepository
}

func NewEventRepository(collection *mongo.Collection) core.EventRepository {
	return &eventRepository{
		collection: collection,
		delegate:   NewMongoDbEventRepository(collection),
	}
}

func (m *eventRepository) FindByStartAfterAndEndBefore(from, to time.Time) ([]*core.Event, error) {
	dbArray, err := m.delegate.FindByStartsAtGreaterThanEqualAndEndsAtLessThanEqual(context.Background(), from, to)
	if err != nil {
		return []*core.Event{}, err
	}
	events, err := MapToEventArray(dbArray)
	if err != nil {
		return []*core.Event{}, err
	}
	return events, nil
}

func (m *eventRepository) Update(toUpdate []*core.Event) error {
	errorList := ""
	for _, event := range toUpdate {
		_, err := m.updateMinPriceAndMaxPrice(context.Background(), event)
		if err != nil {
			errorList += "\n\t" + err.Error()
		}
	}
	if len(errorList) != 0 {
		return errors.NewGenericError("Error updating event list:" + errorList)
	}
	return nil
}

func (m *eventRepository) InsertOrUpdate(toInsert []*core.Event) error {
	errorList := ""
	for _, event := range toInsert {
		ok, err := m.updateMinPriceAndMaxPrice(context.Background(), event)
		if err != nil {
			errorList += "\n\t" + err.Error()
			continue
		}
		if !ok {
			db, err := MapToEventDB(event)
			if err != nil {
				errorList += "\n\t" + err.Error()
				continue
			}
			db.Id = uuid.New()
			if err != nil {
				errorList += "\n\t" + err.Error()
				continue
			}
			_, err = m.delegate.InsertOne(context.Background(), db)
			if err != nil {
				errorList += "\n\t" + err.Error()
			}
		}
	}
	if len(errorList) != 0 {
		return errors.NewGenericError("Error updating event list:" + errorList)
	}
	return nil
}

func (m *eventRepository) updateMinPriceAndMaxPrice(ctx context.Context, event *core.Event) (bool, error) {
	filter := bson.M{"prov_id": event.ProvId}
	update := bson.D{
		{"$min", bson.D{{"min_price", event.MinPrice}}},
		{"$max", bson.D{{"max_price", event.MaxPrice}}},
	}
	result, err := m.collection.UpdateOne(ctx, filter, update)
	return result.MatchedCount != 0, err
}
