package core

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type supplyServiceSuite struct {
	suite.Suite
	eventProvider *MockEventProvider
	eventService  *MockEventService
	supplyService SupplyService
}

func (suite *supplyServiceSuite) SetupSuite() {
	suite.eventProvider = NewMockEventProvider(suite.T())
	suite.eventService = NewMockEventService(suite.T())
	suite.supplyService = NewSupplyService(suite.eventService, suite.eventProvider)
}

func TestSupplyServiceSuite(t *testing.T) {
	suite.Run(t, new(supplyServiceSuite))
}

func (suite *supplyServiceSuite) TestFetchEventsFromProvider_Successful() {
	events := getEventList(6)
	suite.eventProvider.EXPECT().GetEvents().Return(events, nil)
	suite.eventService.EXPECT().UpdateEvents(events)

	suite.supplyService.FetchEventsFromProvider()
}

func (suite *supplyServiceSuite) TestFetchEventsFromProvider_SuccessfulWhenProviderFails() {
	suite.eventProvider.EXPECT().GetEvents().Return([]*Event{}, errors.New("unexpected error"))

	suite.supplyService.FetchEventsFromProvider()
	suite.eventService.AssertNotCalled(suite.T(), "UpdateEvents")
}
