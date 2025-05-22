package controllers

import (
	"github.com/google/uuid"
	"time"
)

type EventSummary struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	StartTime time.Time `json:"start_time"`
	EndDate   time.Time `json:"end_date"`
	EndTime   time.Time `json:"end_time"`
	MinPrice  float64   `json:"min_price"`
	MaxPrice  float64   `json:"max_price"`
}

type EventList struct {
	Events []EventSummary `json:"events"`
}
