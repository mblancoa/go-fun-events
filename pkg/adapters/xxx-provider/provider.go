package xxx_provider

import (
	"encoding/xml"
	core2 "github.com/mblancoa/go-fun-events/pkg/core"
	"github.com/mblancoa/go-fun-events/pkg/errors"
	"gopkg.in/resty.v1"
	"time"
)

type xxxEventProvider struct {
	providerName, fetchEventsUrl string
	timeOut                      time.Duration
}

func NewEventProvider(name, fetchEventsUrl string,
	timeOut time.Duration) core2.EventProvider {
	return &xxxEventProvider{
		providerName:   name,
		fetchEventsUrl: fetchEventsUrl,
		timeOut:        timeOut,
	}
}

func (x *xxxEventProvider) GetEvents() ([]*core2.Event, error) {
	client := resty.New()
	client.SetTimeout(x.timeOut)

	res, err := client.R().Get(x.fetchEventsUrl)
	if err != nil {
		return []*core2.Event{}, errors.NewGenericError("Error fetching events. " + err.Error())
	}

	response := &ProviderResponse{}
	err = xml.Unmarshal(res.Body(), response)
	if err != nil {
		return []*core2.Event{}, errors.NewGenericError("Error unmarshalling response. " + err.Error())
	}

	result, err := MapProviderResponseToEventList(response, x.providerName)
	if err != nil {
		return []*core2.Event{}, errors.NewGenericError("Error mapping response. " + err.Error())
	}
	return result, nil
}
