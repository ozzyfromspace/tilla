package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"tilla/models"
	"time"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// "c_ath9g8cs5n551v6ve3b96m6tio@group.calendar.google.com" -- calendar id for Jack Btesh
const apiKey = "AIzaSyAp7UvmUT4SaeDat1z7dvrT1iFrjTF0res"
const DEFAULT_TEACHER_NAME = "EINSTEIN"
const PaymentStatus = "TO BE INVOICED"

type Calendar struct {
	db *models.Database
}

func NewCalendar(db *models.Database) *Calendar {
	if db == nil {
		log.Fatal("database cannot be nil")
	}

	return &Calendar{db: db}
}

func (cal *Calendar) GetEvents(studentId string, minLocalTime string, maxLocalTime string) (*[]models.Event, error) {
	fmt.Println("getting events")

	returnedStudent, err := cal.db.GetStudentById(studentId)

	if err != nil {
		return nil, err
	}

	log.Println("returned student", returnedStudent)

	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	evts := calendarService.Events
	listCall := evts.List(returnedStudent.CalendarId)

	listCall = listCall.TimeMin(minLocalTime).TimeMax(maxLocalTime).SingleEvents(true)
	e, err := listCall.Do()

	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, v := range e.Items {
		newEvent := models.NewEvent()

		// fmt.Printf("##########>>>>>>>>>>>>> Summary: %+v\nStarting Time: %+v\n", v.Summary, v.Start.DateTime)
		fmt.Printf("####################START %v :::\n", v.Summary)
		// fmt.Println("Course:", v.Summary)
		// fmt.Println("Teacher:", v.Summary)
		// fmt.Println("Student:", v.Summary)
		// fmt.Println("Date:", v.Start.DateTime)
		// fmt.Println("Duration:", v.Start.DateTime, v.End.DateTime)
		// fmt.Println("Fee", "call db for student data...")
		// fmt.Println("Payment Status:", "TO BE INVOICED")
		newEvent.StudentName = strings.ToLower(fmt.Sprintf("%v-%v", returnedStudent.FirstName, returnedStudent.LastName))
		foundTeacher, err := cal.getTeacher(strings.ToLower(v.Summary))

		if err != nil {
			newEvent.Teacher = DEFAULT_TEACHER_NAME
		} else {
			fullName := fmt.Sprintf("%v %v", foundTeacher.FirstName, foundTeacher.LastName)
			newEvent.Teacher = fullName
		}

		timeLayout := "2006-01-02T15:04:05-07:00"
		startTime, _ := time.Parse(timeLayout, v.Start.DateTime)
		endTime, _ := time.Parse(timeLayout, v.End.DateTime)

		fmt.Println("INFO ABOUT TIMES:", v.Start.DateTime, startTime, v.End.DateTime, endTime)
		eventDuration := endTime.Sub(startTime)
		fmt.Println("duration:", eventDuration)

		subject, price, err := cal.getSubjectAndPrice(v.Summary, returnedStudent.Subjects)

		if err != nil {
			log.Println("COULD NOT FIND SUBJECT! DROPPING")
			continue
		}

		hourlyDuration := math.Floor(eventDuration.Hours()*100) / 100

		newEvent.Course = subject
		newEvent.Fee = price * hourlyDuration
		newEvent.Duration = fmt.Sprint(hourlyDuration)
		newEvent.Date = getDateString(&startTime)
		newEvent.PaymentStatus = PaymentStatus

		fmt.Printf("\nEVENT {%v}::: %+v :::EVENT\n\n", v.Summary, newEvent)
		fmt.Println("####################END :::")
	}

	return &[]models.Event{}, nil
}

func (cal *Calendar) getTeacher(summaryStr string) (*models.Teacher, error) {
	if !strings.Contains(summaryStr, " - ") {
		return nil, errors.New("could not find `-` symbol in event summary")
	}

	words := strings.Split(summaryStr, " - ")
	nameSection := words[len(words)-1]

	if !strings.Contains(summaryStr, " & ") {
		return nil, errors.New("could not find `&` symbol in event summary")
	}

	words = strings.Split(nameSection, " & ")
	teacherNickname := words[0]
	return cal.db.GetTeacherByNickname(strings.Trim(teacherNickname, " "))
}

func (cal *Calendar) getSubjectAndPrice(summaryStr string, subjects map[string]float64) (string, float64, error) {
	str := summaryStr

	if strings.Contains(summaryStr, " - ") {
		str = strings.Split(str, " - ")[0]
	}

	str = models.ComputeSubjectName(str)
	fmt.Println("test subject -", str)
	price, ok := subjects[str]

	if !ok {
		return "", 0, errors.New("could not find subject")
	}

	return str, price, nil
}

func getDateString(datetime *time.Time) string {
	year, month, day := datetime.Date()

	return fmt.Sprintf("%02v/%02v/%v", int(month), day, year)
}
