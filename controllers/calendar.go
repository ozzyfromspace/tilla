package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
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

type DroppedEvent struct {
	Summary string
	Date    string
	Student string
}

type Calendar struct {
	db *models.Database
}

func NewCalendar(db *models.Database) *Calendar {
	if db == nil {
		log.Fatal("database cannot be nil")
	}

	return &Calendar{db: db}
}

func (cal *Calendar) GetEvents(returnedStudent *models.Student, minLocalTime string, maxLocalTime string) (*models.Events, *[]DroppedEvent, error) {

	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, nil, err
	}

	evts := calendarService.Events
	listCall := evts.List(returnedStudent.CalendarId)

	listCall = listCall.TimeMin(minLocalTime).TimeMax(maxLocalTime).SingleEvents(true)
	e, err := listCall.Do()

	if err != nil {
		log.Fatalf("%v", err)
	}

	droppedEvents := &[]DroppedEvent{}
	pickedEvents := &models.Events{}

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

		newEvent.DateTime = startTime

		fmt.Println("INFO ABOUT TIMES:", v.Start.DateTime, startTime, v.End.DateTime, endTime)
		eventDuration := endTime.Sub(startTime)
		fmt.Println("duration:", eventDuration)

		subject, price, err := cal.getSubjectAndPrice(v.Summary, returnedStudent.Subjects)

		if err != nil {
			log.Println("COULD NOT FIND SUBJECT! DROPPING")
			*droppedEvents = append(*droppedEvents, DroppedEvent{
				Summary: v.Summary,
				Date:    v.Start.DateTime,
				Student: fmt.Sprintf("%v %v", returnedStudent.FirstName, returnedStudent.LastName),
			})
			continue
		}

		hourlyDuration := math.Floor(eventDuration.Hours()*100) / 100

		newEvent.Course = formatSubject(subject)
		newEvent.Fee = price * hourlyDuration
		newEvent.Duration = fmt.Sprint(hourlyDuration)
		newEvent.Date = getDateString(&startTime)
		newEvent.PaymentStatus = PaymentStatus

		// startTime.
		t := fmt.Sprintf("%02v:%02v", startTime.Hour(), startTime.Minute())
		newEvent.Time = t

		*pickedEvents = append(*pickedEvents, *newEvent)

		fmt.Printf("\nEVENT {%v}::: %+v :::EVENT\n\n", v.Summary, newEvent)
		fmt.Println("####################END :::")
	}

	log.Println("Dropped Events", droppedEvents)
	sort.Sort(pickedEvents)

	return pickedEvents, droppedEvents, nil
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

func formatSubject(str string) string {
	strSlice := strings.Split(str, "_")
	output := []string{}

	for _, s := range strSlice {
		sRune := []rune(s)
		l1 := strings.ToUpper(string(sRune[0]))
		l2 := strings.ToLower(string(sRune[1:]))
		word := fmt.Sprintf("%v%v", l1, l2)
		output = append(output, word)
	}

	return strings.Join(output, " ")
}
