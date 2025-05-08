package xxx_provider

import (
	"bytes"
	"github.com/h2non/gock"
	"github.com/mblanco/Go-Acme-events/core/ports"
	"github.com/mblanco/Go-Acme-events/tools"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var providerName = "xxx"
var providerDomain = "https://xxx-provider.com"
var fetchEventsPath = "/events"
var fetchEventsUrl = providerDomain + fetchEventsPath
var timeout = 2 * time.Millisecond

type xxxProviderSuite struct {
	suite.Suite
	provider ports.EventProvider
}

func (suite *xxxProviderSuite) SetupSuite() {
	suite.provider = NewEventProvider(providerName, fetchEventsUrl, timeout)
}
func (suite *xxxProviderSuite) SetupTest() {
	gock.Clean()
}

func TestXxxProviderSuite(t *testing.T) {
	suite.Run(t, new(xxxProviderSuite))
}

func (suite *xxxProviderSuite) TestGetEvents_successful() {
	defer gock.Off()
	rsp := tools.LoadFile("testdata/response.xml")
	gock.New(providerDomain).Get(fetchEventsPath).Reply(200).Body(bytes.NewReader(rsp))

	events, err := suite.provider.GetEvents()

	suite.Assert().Nil(err)
	suite.Assert().NotEmpty(events)
	suite.Assert().Equal(4, len(events))
}

func (suite *xxxProviderSuite) TestGetEvents_returnsErrorWhenResponseFails() {
	defer gock.Off()
	gock.New(providerDomain).Get(fetchEventsPath).Reply(500).SetError(tools.NewTestError("Internal server error"))

	events, err := suite.provider.GetEvents()

	suite.Assert().Empty(events)
	suite.Assert().Error(err)
	suite.Assert().Equal("Error fetching events. Get \"https://xxx-provider.com/events\": Internal server error", err.Error())
}

func (suite *xxxProviderSuite) TestGetEvents_returnsErrorWhenUnmarshallingFails() {
	defer gock.Off()
	rsp := tools.LoadFile("testdata/response-with-error.xml")
	gock.New(providerDomain).Get(fetchEventsPath).Reply(200).Body(bytes.NewReader(rsp))

	events, err := suite.provider.GetEvents()

	suite.Assert().Empty(events)
	suite.Assert().Error(err)
	suite.Assert().Equal("Error unmarshalling response. XML syntax error on line 10: element <output> closed by </planList>", err.Error())
}

func (suite *xxxProviderSuite) TestGetEvents_returnsErrorWhenMappingFails() {
	defer gock.Off()
	rsp := tools.LoadFile("testdata/response-with-bad-date-format.xml")
	gock.New(providerDomain).Get(fetchEventsPath).Reply(200).Body(bytes.NewReader(rsp))

	events, err := suite.provider.GetEvents()

	suite.Assert().Empty(events)
	suite.Assert().Error(err)
	suite.Assert().Equal("Error mapping response. Error mapping event: parsing time \"2021-06-30 21:00:00\" as \"2006-01-02T15:04:05\": cannot parse \" 21:00:00\" as \"T\"", err.Error())
}
