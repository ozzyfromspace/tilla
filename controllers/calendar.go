package controllers

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// "c_ath9g8cs5n551v6ve3b96m6tio@group.calendar.google.com" -- calendar id for Jack Btesh
var apiKey = "AIzaSyAp7UvmUT4SaeDat1z7dvrT1iFrjTF0res"

type Calendar struct{}

func (cal *Calendar) GetEvents(calendarId string) ([]*calendar.Event, error) {
	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	evts := calendarService.Events
	listCall := evts.List(calendarId)

	listCall = listCall.TimeMin("2023-02-27T07:59:59-05:00").TimeMax("2023-03-03T19:00:01-05:00").SingleEvents(true)
	e, err := listCall.Do()

	if err != nil {
		log.Fatalf("%v", err)
	}

	for _, v := range e.Items {
		fmt.Printf("##########>>>>>>>>>>>>> Summary: %+v\nStarting Time: %+v\n", v.Summary, v.Start.DateTime)
	}

	return []*calendar.Event{}, nil
}
