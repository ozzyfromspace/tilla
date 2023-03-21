package models

type SubjectMetaData struct{}

type Subject struct {
	Name                   string  `json:"name"`
	PricePerSession        float64 `json:"pricePerSession" bson:"pricePerSession,required"`
	SessionLengthInMinutes int     `json:"sessionLength" bson:"sessionLength,required"`
}

type SessionData struct {
	PricePerSession        float64
	SessionLengthInMinutes int
	OriginalFilename       string
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
