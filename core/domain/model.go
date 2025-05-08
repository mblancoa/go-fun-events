package domain

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id           uuid.UUID
	ProvId       string
	Title        string
	IsOnlineSale bool
	StartsAt     time.Time
	EndsAt       time.Time
	MaxPrice     float64
	MinPrice     float64
}
