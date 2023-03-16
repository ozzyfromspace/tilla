package models

type TeacherMetaData struct{}

func NewTeacher() *Teacher {
	return &Teacher{}
}

func NewTeachers() *[]Teacher {
	return &[]Teacher{}
}

type Teacher struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Nickname  string `json:"nickname" bson:"nickname"`
}
