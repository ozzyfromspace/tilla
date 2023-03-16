package handler

import (
	"fmt"
	"net/http"
	"tilla/apis"
	"tilla/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go!")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	db := models.NewDatabase()

	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	app.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	apis.NewStudentApi(app, db).RegisterApi()
	apis.NewTeacherApi(app, db).RegisterApi()
	apis.NewCalendarApi(app, db).RegisterApi()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, models.MsgPayload("route not implemented"))
	})

	// app.Run(":8080")
	app.ServeHTTP(w, r)
}
