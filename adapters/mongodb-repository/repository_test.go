package mongodb_repository

import (
	"context"
	"github.com/google/uuid"
	core2 "github.com/mblancoa/go-fun-events/pkg/core"
	"github.com/mblancoa/go-fun-events/tools"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type eventRepositorySuite struct {
	mongoDBSuite
	eventRepository core2.EventRepository
}

func (suite *eventRepositorySuite) SetupSuite() {
	suite.mongoDBSuite.SetupSuite()
	suite.setupEventsCollection()
	suite.eventRepository = NewEventRepository(suite.eventsCollection)
}

func (suite *eventRepositorySuite) SetupTest() {
	ctx := context.Background()
	deleteAll(suite.eventsCollection, ctx)
}

func (suite *eventRepositorySuite) TearDownSuite() {
	suite.mongoDBSuite.TearDownSuite()
}

func TestCredentialsServiceSuite(t *testing.T) {
	suite.Run(t, new(eventRepositorySuite))
}

func (suite *eventRepositorySuite) TestFindByStartAfterAndEndBefore_successful() {
	list := getEventDBList(4)

	from := time.Now()
	to := from.Add(24 * time.Hour)

	for i, db := range list {
		db.StartsAt = from.Add(-time.Duration(i) * time.Hour)
		db.EndsAt = to.Add(time.Duration(i) * time.Hour)
		insertOne(suite.eventsCollection, context.TODO(), db)
	}

	events, err := suite.eventRepository.FindByStartAfterAndEndBefore(from, to)

	suite.Assert().Empty(err)
	suite.Assert().NotEmpty(events)
	suite.Assert().Equal(1, len(events))
}

func (suite *eventRepositorySuite) TestUpdate() {
	list := getEventDBList(4)

	from := time.Now()
	to := from.Add(24 * time.Hour)

	for i, _ := range list {
		list[i].StartsAt = from.Add(-time.Duration(i) * time.Hour)
		list[i].EndsAt = to.Add(time.Duration(i) * time.Hour)
		list[i].MinPrice = 40
		list[i].MaxPrice = 40
		insertOne(suite.eventsCollection, context.TODO(), list[i])
	}

	event, _ := MapToEvent(&list[0])
	event.MinPrice = 45
	event.MaxPrice = 65
	event1, _ := MapToEvent(&list[1])
	event1.MinPrice = 12.5
	event1.MaxPrice = 30

	err := suite.eventRepository.Update([]*core2.Event{event, event1})

	suite.Assert().Empty(err)
	db := &EventDB{}
	findOne(suite.eventsCollection, context.TODO(), "_id", event.Id, &db)
	suite.Assert().NotEmpty(db)
	suite.Assert().Equal(40.0, db.MinPrice)
	suite.Assert().Equal(event.MaxPrice, db.MaxPrice)

	db1 := &EventDB{}
	findOne(suite.eventsCollection, context.TODO(), "_id", event1.Id, &db1)
	suite.Assert().NotEmpty(db1)
	suite.Assert().Equal(event1.MinPrice, db1.MinPrice)
	suite.Assert().Equal(40.0, db1.MaxPrice)
}

func (suite *eventRepositorySuite) TestInsertOrUpdate() {
	start := time.Now()
	end := start.Add(time.Hour)
	list := []EventDB{
		{Id: uuid.New(), ProvId: "xxx-001", Title: "title1", StartsAt: start, EndsAt: end, MinPrice: 15.0, MaxPrice: 30.0},
		{Id: uuid.New(), ProvId: "xxx-002", Title: "title1", StartsAt: start.Add(24 * time.Hour), EndsAt: end.Add(24 * time.Hour), MinPrice: 15.0, MaxPrice: 30.0},
	}
	for _, v := range list {
		insertOne(suite.eventsCollection, context.Background(), v)
	}
	suite.Assert().Equal(int64(2), count(suite.eventsCollection, context.Background()))
	db := &EventDB{}
	findOne(suite.eventsCollection, context.TODO(), "prov_id", "xxx-001", &db)
	suite.Assert().NotEmpty(db)
	suite.Assert().Equal(15.0, db.MinPrice)
	suite.Assert().Equal(30.0, db.MaxPrice)

	newList := []*core2.Event{
		{ProvId: "xxx-001", Title: "title1", StartsAt: start, EndsAt: end, MinPrice: 15.0, MaxPrice: 40.0},
		{ProvId: "xxx-003", Title: "title1", IsOnlineSale: false, StartsAt: start.Add(24 * time.Hour), EndsAt: end.Add(24 * time.Hour), MinPrice: 15.0, MaxPrice: 30.0},
	}

	err := suite.eventRepository.InsertOrUpdate(newList)

	suite.Assert().Empty(err)
	count := count(suite.eventsCollection, context.Background())
	suite.Assert().Equal(int64(3), count)

	db1 := EventDB{}
	findOne(suite.eventsCollection, context.TODO(), "prov_id", "xxx-001", &db1)
	suite.Assert().NotEmpty(db1)
	suite.Assert().Equal(15.0, db1.MinPrice)
	suite.Assert().Equal(40.0, db1.MaxPrice)

	db3 := EventDB{}
	found := findOne(suite.eventsCollection, context.TODO(), "prov_id", "xxx-003", &db3)
	suite.Assert().NotEmpty(db3)
	suite.Assert().True(found)
}

func getEventDBList(size int) []EventDB {
	var list []EventDB
	for i := 0; i < size; i++ {
		db := EventDB{}
		tools.FakerBuild(&db)
		db.Id = uuid.New()
		list = append(list, db)
	}
	return list
}
