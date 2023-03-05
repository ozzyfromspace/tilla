package main

import (
	"fmt"
	"log"
	"tilla/controllers"
	"tilla/models"
	"time"
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

	timeLayout := "2006-01-02T15:04:05-07:00"
	t, _ := time.Parse(timeLayout, minLocalTime)

	cal := controllers.NewCalendar(db)
	fmt.Println("getting events")

	returnedStudent, err := db.GetStudentById(studentId)

	if err != nil {
		log.Fatal("could not find student!")
	}

	log.Println("returned student", returnedStudent)
	calEvents, droppedEvents, err := cal.GetEvents(returnedStudent, minLocalTime, maxLocalTime)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("CHOSEN EVENTS -", calEvents)
	log.Println("DROPPED EVENTS -", droppedEvents)

	convDoc := &[]controllers.ConversionDoc{
		{Student: returnedStudent, Events: calEvents},
	}

	controllers.GenerateExcel(convDoc, int(t.Month()), t.Year())
}
