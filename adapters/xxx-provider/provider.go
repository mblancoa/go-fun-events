package xxx_provider

import (
	"encoding/xml"
	"github.com/mblanco/go-fun-events/core"
	"github.com/mblanco/go-fun-events/errors"
	"gopkg.in/resty.v1"
	"time"
)

type xxxEventProvider struct {
	providerName, fetchEventsUrl string
	timeOut                      time.Duration
}

func NewEventProvider(name, fetchEventsUrl string,
	timeOut time.Duration) core.EventProvider {
	return &xxxEventProvider{
		providerName:   name,
		fetchEventsUrl: fetchEventsUrl,
		timeOut:        timeOut,
	}
}

func (x *xxxEventProvider) GetEvents() ([]*core.Event, error) {
	client := resty.New()
	client.SetTimeout(x.timeOut)

	res, err := client.R().Get(x.fetchEventsUrl)
	if err != nil {
		return []*core.Event{}, errors.NewGenericError("Error fetching events. " + err.Error())
	}

	response := &ProviderResponse{}
	err = xml.Unmarshal(res.Body(), response)
	if err != nil {
		return []*core.Event{}, errors.NewGenericError("Error unmarshalling response. " + err.Error())
	}

	result, err := MapProviderResponseToEventList(response, x.providerName)
	if err != nil {
		return []*core.Event{}, errors.NewGenericError("Error mapping response. " + err.Error())
	}
	return result, nil
}
