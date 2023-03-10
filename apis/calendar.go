package apis

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	api.downloadExcel()
}

type LocalTimes struct {
	MinLocalTime string `json:"minLocalTime"`
	MaxLocalTime string `json:"maxLocalTime"`
}

func (lt *LocalTimes) sort() error {
	timeLayout := "2006-01-02T15:04:05-07:00"
	tmin, err := time.Parse(timeLayout, lt.MinLocalTime)

	if err != nil {
		alternativeTimeLayout := "2006-01-02T15:04:05Z"
		tmin, err = time.Parse(alternativeTimeLayout, lt.MinLocalTime)

		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}

	tmax, err := time.Parse(timeLayout, lt.MaxLocalTime)

	if err != nil {
		alternativeTimeLayout := "2006-01-02T15:04:05Z"
		tmax, err = time.Parse(alternativeTimeLayout, lt.MaxLocalTime)

		if err != nil {
			return err
		}
	}

	if tmax.Before(tmin) {
		lt.MaxLocalTime, lt.MinLocalTime = lt.MinLocalTime, lt.MaxLocalTime
	}

	return nil
}

func (api *CalendarApi) createExcel() {
	api.r.POST("/excel", func(c *gin.Context) {
		localTimes := &LocalTimes{}
		if err := c.BindJSON(localTimes); err != nil {
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("could not process one or more local times"))
			return
		}

		localTimes.sort()
		cal := controllers.NewCalendar(api.db)
		filepath, droppedEventsMap, err := cal.ToExcel(localTimes.MinLocalTime, localTimes.MaxLocalTime)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("failed to create excel file"))
			return
		}

		for k, v := range droppedEventsMap {
			fmt.Printf("Student - %v (%v):\n\n", (*v)[0].Student, k)

			if *v == nil {
				continue
			}

			for i, v2 := range *v {
				y, m, d := v2.Date.Date()
				datetimeStr := fmt.Sprintf("%v %02d, %04v starting at %02d:%02d", m, d, y, v2.Date.Hour(), v2.Date.Minute())
				fmt.Printf("\t%5d. %v on %v\n\n", i+1, v2.Summary, datetimeStr)
			}

			fmt.Printf("\n\n")
		}

		c.JSON(http.StatusCreated, gin.H{
			"msg":      "excel created!",
			"filename": filepath,
		})
	})
}

func (api *CalendarApi) downloadExcel() {
	api.r.GET("/excel/:filename", func(c *gin.Context) {
		filename, found := c.Params.Get("filename")

		if !found {
			c.JSON(http.StatusBadRequest, models.MsgPayload("excel filename was not provided"))
			return
		}

		file, err := os.Open(filename)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusBadRequest, models.MsgPayload("could not find excel file"))
			return
		}

		fileinfo, err := file.Stat()

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve created excel file stats"))
			return
		}

		contentLength := fileinfo.Size()
		contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
		contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", fileinfo.Name())

		extraHeaders := map[string]string{
			"Content-Disposition": contentDisposition,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, file, extraHeaders)
	})
}
