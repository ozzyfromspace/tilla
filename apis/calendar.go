package apis

import (
	"fmt"
	"log"
	"net/http"
	"tilla/controllers"
	"tilla/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CalendarApi struct {
	r  *gin.Engine
	db *models.Database
}

func NewCalendarApi(r *gin.Engine, db *models.Database) *CalendarApi {
	if r == nil {
		log.Fatal("gin engine cannot be nil")
	}

	if db == nil {
		log.Fatal("database cannot be nil")
	}

	return &CalendarApi{r: r, db: db}
}

func (api *CalendarApi) RegisterApi() {
	api.createExcel()
}

type LocalTimes struct {
	MinLocalTime string `json:"minLocalTime"`
	MaxLocalTime string `json:"maxLocalTime"`
}

func (lt *LocalTimes) sort() error {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tmin, err := time.Parse(timeLayout, lt.MinLocalTime)

	if err != nil {
		return err
	}

	tmax, err := time.Parse(timeLayout, lt.MaxLocalTime)

	if err != nil {
		return err
	}

	if tmax.Before(tmin) {
		lt.MaxLocalTime, lt.MinLocalTime = lt.MinLocalTime, lt.MaxLocalTime
	}

	return nil
}

func (api *CalendarApi) createExcel() {
	api.r.POST("/students/excel", func(c *gin.Context) {
		localTimes := &LocalTimes{}
		if err := c.BindJSON(localTimes); err != nil {
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("could not process one or more local times"))
		}

		localTimes.sort()
		cal := controllers.NewCalendar(api.db)
		droppedEventsMap, err := cal.ToExcel(localTimes.MinLocalTime, localTimes.MaxLocalTime)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("failed to create excel file"))
		}

		// logs
		for k, v := range droppedEventsMap {
			fmt.Printf("Student - %v (%v):\n\n", (*v)[0].Student, k)

			for i, v2 := range *v {
				y, m, d := v2.Date.Date()
				datetimeStr := fmt.Sprintf("%v %02d, %04v starting at %02d:%02d", m, d, y, v2.Date.Hour(), v2.Date.Minute())
				fmt.Printf("\t%5d. %v on %v\n\n", i+1, v2.Summary, datetimeStr)
			}

			fmt.Printf("\n\n")
		}

		c.JSON(http.StatusCreated, models.MsgPayload("excel file has been created"))
	})

}
