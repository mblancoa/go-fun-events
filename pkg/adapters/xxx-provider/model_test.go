package xxx_provider

import (
	tools2 "github.com/mblancoa/go-fun-events/pkg/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapProviderResponseToEventList_successful(t *testing.T) {
	rsp := &ProviderResponse{}
	err := tools2.UnmarshalXmlResource("testdata/to-map.xml", rsp)
	tools2.ManageTestError(err)

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.NotEmpty(t, eventsA)
	assert.Equal(t, 1, len(eventsA))
	event := eventsA[0]
	assert.NotEmpty(t, event)
	assert.Equal(t, "xxx-291-291", event.ProvId)
	assert.Equal(t, "Camela en concierto", event.Title)
	assert.True(t, event.IsOnlineSale)
	startsAt := "2021-06-30T21:00:00"
	assert.Equal(t, startsAt, event.StartsAt.Format(DateTimeLayout))
	endsAt := "2021-06-30T22:00:00"
	assert.Equal(t, endsAt, event.EndsAt.Format(DateTimeLayout))
	assert.Equal(t, 15.00, event.MinPrice)
	assert.Equal(t, 20.00, event.MaxPrice)
}
func TestMapProviderResponseToEventList_successfulPricesAreOIfNotZones(t *testing.T) {
	rsp := &ProviderResponse{}
	err := tools2.UnmarshalXmlResource("testdata/to-map-without-zones.xml", rsp)
	tools2.ManageTestError(err)

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.NotEmpty(t, eventsA)
	assert.Equal(t, 1, len(eventsA))
	event := eventsA[0]
	assert.NotEmpty(t, event)
	assert.Equal(t, "xxx-291-291", event.ProvId)
	assert.Equal(t, "Camela en concierto", event.Title)
	assert.True(t, event.IsOnlineSale)
	startsAt := "2021-06-30T21:00:00"
	assert.Equal(t, startsAt, event.StartsAt.Format(DateTimeLayout))
	endsAt := "2021-06-30T22:00:00"
	assert.Equal(t, endsAt, event.EndsAt.Format(DateTimeLayout))
	assert.Equal(t, 0.00, event.MinPrice)
	assert.Equal(t, 0.00, event.MaxPrice)
}

func TestMapProviderResponseToEventList_returnsEmptyArrayWhenResponseIsEmpty(t *testing.T) {
	rsp := &ProviderResponse{}

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.Empty(t, eventsA)
}

func TestMapProviderResponseToEventList_returnsEmptyArrayWhenResponseOutputIsEmpty(t *testing.T) {
	rsp := &ProviderResponse{Output: Output{}}

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.Empty(t, eventsA)
}

func TestMapProviderResponseToEventList_returnsEmptyArrayWhenResponseOutputBasePlansIsEmpty(t *testing.T) {
	rsp := &ProviderResponse{Output: Output{BasePlans: []BasePlan{}}}

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.Empty(t, eventsA)
}

func TestMapProviderResponseToEventList_doesNotMapWhenPlansIsEmpty(t *testing.T) {
	rsp := &ProviderResponse{Output: Output{BasePlans: []BasePlan{{Plans: []Plan{}}}}}

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.Empty(t, eventsA)
}

func TestMapProviderResponseToEventList_doesNotMapWhenPlansIsNil(t *testing.T) {
	rsp := &ProviderResponse{Output: Output{BasePlans: []BasePlan{{SellMode: "online"}}}}

	eventsA, err := MapProviderResponseToEventList(rsp, providerName)

	assert.Nil(t, err)
	assert.Empty(t, eventsA)
}
