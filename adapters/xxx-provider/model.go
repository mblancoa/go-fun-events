package xxx_provider

import (
	"encoding/xml"
	"fmt"
	"github.com/mblanco/Go-Acme-events/core/domain"
	"github.com/mblanco/Go-Acme-events/errors"
	"math"
	"reflect"
	"time"
)

const (
	ProviderIdPattern = "%s-%s-%s"
	DateTimeLayout    = "2006-01-02T15:04:05"
)

type ProviderResponse struct {
	XMLName xml.Name `xml:"planList"`
	Output  Output   `xml:"output"`
}

type Output struct {
	BasePlans []BasePlan `xml:"base_plan"`
}

type BasePlan struct {
	XMLName            xml.Name `xml:"base_plan"`
	BasePlanID         string   `xml:"base_plan_id,attr"`
	SellMode           string   `xml:"sell_mode,attr"`
	Title              string   `xml:"title,attr"`
	OrganizerCompanyID string   `xml:"organizer_company_id,attr,omitempty"`
	Plans              []Plan   `xml:"plan"`
}

type Plan struct {
	XMLName       xml.Name `xml:"plan"`
	PlanStartDate string   `xml:"plan_start_date,attr"`
	PlanEndDate   string   `xml:"plan_end_date,attr"`
	PlanID        string   `xml:"plan_id,attr"`
	SellFrom      string   `xml:"sell_from,attr"`
	SellTo        string   `xml:"sell_to,attr"`
	SoldOut       string   `xml:"sold_out,attr"`
	Zones         []Zone   `xml:"zone"`
}

type Zone struct {
	XMLName  xml.Name `xml:"zone"`
	ZoneID   string   `xml:"zone_id,attr"`
	Capacity string   `xml:"capacity,attr"`
	Price    float64  `xml:"price,attr"`
	Name     string   `xml:"providerName,attr"`
	Numbered string   `xml:"numbered,attr"`
}

func MapProviderResponseToEventList(response *ProviderResponse, providerName string) ([]*domain.Event, error) {
	if reflect.DeepEqual(response, ProviderResponse{}) {
		return []*domain.Event{}, nil
	}
	if reflect.DeepEqual(response.Output, Output{}) {
		return []*domain.Event{}, nil
	}

	var eventsA []*domain.Event
	for _, basePlan := range response.Output.BasePlans {
		for _, plan := range basePlan.Plans {
			event := &domain.Event{}
			event.ProvId = fmt.Sprintf(ProviderIdPattern, providerName, basePlan.BasePlanID, plan.PlanID)
			event.Title = basePlan.Title
			event.IsOnlineSale = "online" == basePlan.SellMode
			startsAt, err := time.Parse(DateTimeLayout, plan.PlanStartDate)
			if err != nil {
				return []*domain.Event{}, errors.NewGenericError("Error mapping event: " + err.Error())
			}
			event.StartsAt = startsAt
			endsAt, err := time.Parse(DateTimeLayout, plan.PlanEndDate)
			if err != nil {
				return []*domain.Event{}, errors.NewGenericError("Error mapping event: " + err.Error())
			}
			event.EndsAt = endsAt
			if plan.Zones != nil && len(plan.Zones) != 0 {
				event.MinPrice = math.MaxFloat64
				for _, zone := range plan.Zones {
					if zone.Price < event.MinPrice {
						event.MinPrice = zone.Price
					}
					if zone.Price > event.MaxPrice {
						event.MaxPrice = zone.Price
					}
				}
			}
			eventsA = append(eventsA, event)
		}
	}

	return eventsA, nil
}
