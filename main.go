package main

import (
	"tilla/apis"
	"tilla/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := models.NewDatabase()

	r := gin.Default()

	apis.NewStudentApi(r, db).RegisterApi()
	apis.NewCalendarApi(r, db).RegisterApi()
	apis.NewTeacherApi(r, db).RegisterApi()

	r.Run(":8080")
}
