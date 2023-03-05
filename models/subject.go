package models

type SubjectMetaData struct{}

type Subject struct {
	Name         string  `json:"name"`
	PricePerHour float64 `json:"pricePerHour" bson:"price,required"`
}

type SubjectPayload struct {
	Subjects  []Subject `json:"subjects"`
	StudentId string    `json:"studentId"`
}

func NewSubjectPayload() *SubjectPayload {
	return &SubjectPayload{}
}

func NewSubject() *Subject {
	return &Subject{}
}

func NewSubjects() *[]Subject {
	return &[]Subject{}
}
