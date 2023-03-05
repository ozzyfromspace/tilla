package main

import (
	"log"
	"tilla/controllers"
	"tilla/models"
)

func main() {
	db := models.NewDatabase()
	//
	// 	r := gin.Default()
	//
	// 	apis.NewStudentApi(r, db).RegisterApi()
	// 	apis.NewTeacherApi(r, db).RegisterApi()
	//
	// 	r.Run(":8080")

	// calId := "c_ath9g8cs5n551v6ve3b96m6tio@group.calendar.google.com"

	studentId := "640443353d1ffb09ae445fdc"

	minLocalTime := "2023-02-27T07:59:59-05:00"
	maxLocalTime := "2023-03-03T19:00:01-05:00"

	cal := controllers.NewCalendar(db)
	calEvents, err := cal.GetEvents(studentId, minLocalTime, maxLocalTime)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("EVENTS", calEvents)
}
