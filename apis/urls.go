package apis

import "fmt"

const NextBaseURL = "http://localhost:3000"

func extendBase(route string) string {
	return fmt.Sprintf("%v%v", NextBaseURL, route)
}

var (
	RStudent            = "/student"
	RStudent_id         = "/student/:id"
	RStudents           = "/students"
	RStudent_subjects   = "/student/subjects"
	RTeacher            = "/teacher"
	RTeacher_id         = "/teacher/:id"
	RTeachers           = "/teachers"
	RCreate_excel       = "/excel"
	RGet_excel          = "/excel/:filename"
	URLStudent          = extendBase("/student")
	URLStudent_id       = extendBase("/student/:id")
	URLStudents         = extendBase("/students")
	URLStudent_subjects = extendBase("/student/subjects")
	URLTeacher          = extendBase("/teacher")
	URLTeacher_id       = extendBase("/teacher/:id")
	URLTeachers         = extendBase("/teachers")
	URLCreate_excel     = extendBase("/excel")
	URLGet_excel        = extendBase("/excel/:filename")
)
