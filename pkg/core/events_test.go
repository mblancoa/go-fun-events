package core

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mblancoa/go-fun-events/pkg/tools"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type eventServiceSuite struct {
	suite.Suite
	eventRepository *MockEventRepository
	eventService    EventService
}

func (suite *eventServiceSuite) SetupSuite() {
	suite.eventRepository = NewMockEventRepository(suite.T())
	suite.eventService = NewEventService(suite.eventRepository)
}

func TestEventServiceSuite(t *testing.T) {
	suite.Run(t, new(eventServiceSuite))
}

func (suite *eventServiceSuite) TestGetEvents_Successful() {
	from := time.Now()
	to := from.Add(time.Hour)
	list := getEventList(4)
	suite.eventRepository.EXPECT().FindByStartAfterAndEndBefore(from, to).Return(list, nil)

	events, err := suite.eventService.GetEvents(from, to)

	suite.Assert().Empty(err)
	suite.Assert().NotEmpty(events)
	suite.Assert().Equal(4, len(events))
}

func (suite *eventServiceSuite) TestGetEvents_ReturnsErrorWhenRepositoryFails() {
	from := time.Now()
	to := from.Add(time.Hour)

	suite.eventRepository.EXPECT().FindByStartAfterAndEndBefore(from, to).Return([]*Event{}, errors.New("unexpected error"))

	events, err := suite.eventService.GetEvents(from, to)

	suite.Assert().Empty(events)
	suite.Assert().Error(err)
	suite.Assert().Equal("unexpected error", err.Error())
}
func (suite *eventServiceSuite) TestUpdateEvents_Successful() {
	list := getEventList(4)
	list[0].IsOnlineSale = false
	list[1].IsOnlineSale = true

	suite.eventRepository.EXPECT().InsertOrUpdate(mock.Anything).Return(nil)
	suite.eventRepository.EXPECT().Update(mock.Anything).Return(nil)

	suite.eventService.UpdateEvents(list)
}

func (suite *eventServiceSuite) TestUpdateEvents_SuccessfulWhenInsertOrUpdateFails() {
	list := getEventList(4)
	list[0].IsOnlineSale = false
	list[1].IsOnlineSale = true

	suite.eventRepository.EXPECT().InsertOrUpdate(mock.Anything).Return(errors.New("unexpected error"))
	suite.eventRepository.EXPECT().Update(mock.Anything).Return(nil)

	suite.eventService.UpdateEvents(list)
}

func (suite *eventServiceSuite) TestUpdateEvents_SuccessfulWhenUpdateFails() {
	list := getEventList(4)
	list[0].IsOnlineSale = false
	list[1].IsOnlineSale = true

	suite.eventRepository.EXPECT().InsertOrUpdate(mock.Anything).Return(nil)
	suite.eventRepository.EXPECT().Update(mock.Anything).Return(errors.New("unexpected error"))

	suite.eventService.UpdateEvents(list)
}

func getEventList(size int) []*Event {
	var list []*Event
	for i := 0; i < size; i++ {
		event := &Event{}
		tools.FakerBuild(event)
		event.Id = uuid.New()
		list = append(list, event)
	}
	return list
}
