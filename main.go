package main

import (
	"net/http"
	"tilla/apis"
	"tilla/models"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	db := models.NewDatabase()

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	apis.NewStudentApi(r, db).RegisterApi()
	apis.NewTeacherApi(r, db).RegisterApi()
	apis.NewCalendarApi(r, db).RegisterApi()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, models.MsgPayload("route not implemented"))
	})
	r.Run(":8080")
}
