package models

type Event struct {
	StudentName   string
	Course        string
	Teacher       string
	Date          string
	Duration      string
	Fee           float64
	PaymentStatus string
}

func NewEvent() *Event {
	return &Event{}
}
