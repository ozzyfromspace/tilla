package models

type Student struct {
	Id         string                 `json:"id" bson:"_id,omitempty"`
	FirstName  string                 `json:"firstName" bson:"firstName"`
	LastName   string                 `json:"lastName" bson:"lastName"`
	Subjects   map[string]SessionData `json:"subjects" bson:"subjects"`
	Nickname   string                 `json:"nickname" bson:"nickname"`
	CalendarId string                 `json:"calendarId" bson:"calendarId,required"`
}

func NewStudent() *Student {
	return &Student{
		Subjects: make(map[string]SessionData),
	}
}

func NewStudents() *[]Student {
	return &[]Student{}
}

type StudentMetaData struct{}
