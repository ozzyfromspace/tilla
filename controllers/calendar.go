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

const apiKey = "AIzaSyAp7UvmUT4SaeDat1z7dvrT1iFrjTF0res"
const DEFAULT_TEACHER_NAME = "EINSTEIN"
const PaymentStatus = "TO BE INVOICED"

type DroppedEvent struct {
	Summary string
	Date    time.Time
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
	e, err := evts.List(returnedStudent.CalendarId).TimeMin(minLocalTime).TimeMax(maxLocalTime).SingleEvents(true).Do()

	if err != nil {
		return nil, nil, err
	}

	droppedEvents := &[]DroppedEvent{}
	pickedEvents := &models.Events{}

	for _, v := range e.Items {
		newEvent := models.NewEvent()
		newEvent.StudentName = strings.ToLower(fmt.Sprintf("%v-%v", returnedStudent.FirstName, returnedStudent.LastName))
		foundTeacher, err := cal.getTeacher(strings.ToLower(v.Summary))

		if err != nil {
			newEvent.Teacher = DEFAULT_TEACHER_NAME
		} else {
			fullName := fmt.Sprintf("%v %v", foundTeacher.FirstName, foundTeacher.LastName)
			newEvent.Teacher = fullName
		}

		timeLayout := "2006-01-02T15:04:05-07:00"
		startTime, err := time.Parse(timeLayout, v.Start.DateTime)

		if err != nil {
			alternativeTimeLayout := "2006-01-02T15:04:05Z"
			startTime, err = time.Parse(alternativeTimeLayout, minLocalTime)

			if err != nil {
				return nil, nil, err
			}
		}

		endTime, err := time.Parse(timeLayout, v.End.DateTime)

		if err != nil {
			alternativeTimeLayout := "2006-01-02T15:04:05Z"
			endTime, err = time.Parse(alternativeTimeLayout, minLocalTime)

			if err != nil {
				return nil, nil, err
			}
		}

		newEvent.DateTime = startTime
		eventDuration := endTime.Sub(startTime)
		subjectData, err := cal.getSubjectData(v.Summary, returnedStudent.Subjects)

		if err != nil {
			*droppedEvents = append(*droppedEvents, DroppedEvent{
				Summary: v.Summary,
				Date:    startTime,
				Student: fmt.Sprintf("%v %v", returnedStudent.FirstName, returnedStudent.LastName),
			})
			continue
		}

		sessionDuration := math.Floor(eventDuration.Minutes()*100) / 100

		_safeSessionLength := subjectData.SessionLengthInMinutes

		if subjectData.SessionLengthInMinutes <= 0 {
			_safeSessionLength = 60
		}

		newEvent.Course = formatSubject(subjectData.SubjectName)
		newEvent.Fee = subjectData.PricePerSession * sessionDuration / float64(_safeSessionLength)
		newEvent.Duration = fmt.Sprint(sessionDuration)
		newEvent.Rate = subjectData.PricePerSession
		newEvent.SessionLengthInMinutes = subjectData.SessionLengthInMinutes
		newEvent.Date = getDateString(&startTime)
		newEvent.PaymentStatus = PaymentStatus

		t := fmt.Sprintf("%02v:%02v", startTime.Hour(), startTime.Minute())
		newEvent.Time = t

		*pickedEvents = append(*pickedEvents, *newEvent)
	}

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

type subjectDataResponse struct {
	SubjectName            string
	PricePerSession        float64
	SessionLengthInMinutes int
}

func (cal *Calendar) getSubjectData(summaryStr string, subjects map[string]models.SessionData) (subjectDataResponse, error) {
	str := summaryStr

	if strings.Contains(summaryStr, " - ") {
		str = strings.Split(str, " - ")[0]
	}

	str = models.ComputeSubjectName(str)
	sessionData, ok := subjects[str]

	if !ok {
		// return subjectDataResponse{}, errors.New("could not find subject")
		return subjectDataResponse{"", 0, 0}, errors.New("could not find subject")
	}

	// return str, price, nil
	return subjectDataResponse{SubjectName: str, PricePerSession: sessionData.PricePerSession, SessionLengthInMinutes: sessionData.SessionLengthInMinutes}, nil
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

func (cal *Calendar) ToExcel(minLocalTime, maxLocalTime string) (string, map[string]*[]DroppedEvent, error) {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tmin, err := time.Parse(timeLayout, minLocalTime)

	if err != nil {
		alternativeTimeLayout := "2006-01-02T15:04:05Z"
		tmin, err = time.Parse(alternativeTimeLayout, minLocalTime)

		if err != nil {
			return "", nil, err
		}
	}
	tmax, err := time.Parse(timeLayout, maxLocalTime)

	if err != nil {
		alternativeTimeLayout := "2006-01-02T15:04:05Z"
		tmax, err = time.Parse(alternativeTimeLayout, maxLocalTime)

		if err != nil {
			return "", nil, err
		}
	}

	students, err := cal.db.GetStudents()

	if err != nil {
		log.Println("could not retrieve students from database")
		return "", nil, err
	}

	convDoc := []ConversionDoc{}
	droppedEventsMap := make(map[string](*[]DroppedEvent))

	for _, s := range *students {
		calEvents, droppedEvents, err := cal.GetEvents(&s, minLocalTime, maxLocalTime)

		if err != nil {
			log.Printf("could not extract events from %v %v's calendar with calendarId %v...\n%v\n", s.FirstName, s.LastName, truncateStr(s.CalendarId), err)
			continue
		}

		key := fmt.Sprintf("%v_%v", s.Nickname, truncateStr(s.CalendarId))
		droppedEventsMap[key] = droppedEvents

		newConversionDoc := ConversionDoc{
			Student: s,
			Events:  *calEvents,
		}

		convDoc = append(convDoc, newConversionDoc)
	}

	filepath, err := generateExcel(&convDoc, tmin.Day(), tmin.Month(), tmin.Year(), tmax.Day(), tmax.Month(), tmax.Year())

	return filepath, droppedEventsMap, err
}
