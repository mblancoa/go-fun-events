package domain

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id         uuid.UUID
	ProvId     string
	title      string
	OnlineSale bool
	StartsAt   time.Time
	EndsAt     time.Time
	MaxPrice   float32
	MinPrice   float32
}
