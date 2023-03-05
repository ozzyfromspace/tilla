package models

import "time"

type Event struct {
	StudentName   string
	Course        string
	Teacher       string
	Date          time.Time
	Duration      string
	Fee           float64
	PaymentStatus string
}

func NewEvent() *Event {
	return &Event{}
}
