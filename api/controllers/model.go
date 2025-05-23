package controllers

import (
	"github.com/google/uuid"
	"time"
)

type EventDto struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	StartTime time.Time `json:"start_time"`
	EndDate   time.Time `json:"end_date"`
	EndTime   time.Time `json:"end_time"`
	MinPrice  float64   `json:"min_price"`
	MaxPrice  float64   `json:"max_price"`
} //@Name EventSummary

type EventList struct {
	Events []EventDto `json:"events"`
} //@Name EventList
