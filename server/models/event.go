package models

import "time"

type Event struct {
	StudentName            string
	Course                 string
	Teacher                string
	Date                   string
	DateTime               time.Time
	Rate                   float64
	SessionLengthInMinutes int
	Time                   string
	Duration               string
	Fee                    float64
	PaymentStatus          string
}

type Events []Event

func NewEvent() *Event {
	return &Event{}
}

func (es *Events) Len() int {
	return len(*es)
}

func (es *Events) Less(i, j int) bool {
	ei := (*es)[i]
	ej := (*es)[j]
	return ej.DateTime.After(ei.DateTime)
}

func (es *Events) Swap(i, j int) {
	(*es)[i], (*es)[j] = (*es)[j], (*es)[i]
}
