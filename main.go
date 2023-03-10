package main

import (
	"net/http"
	"tilla/apis"
	"tilla/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	db := models.NewDatabase()

	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// 	corsConfig := cors.DefaultConfig()
	//
	// 	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// 	// To be able to send tokens to the server.
	// 	corsConfig.AllowCredentials = true
	//
	// 	// OPTIONS method for ReactJS
	// 	corsConfig.AddAllowMethods("OPTIONS")
	//
	// 	// Register the middleware
	// 	r.Use(cors.New(corsConfig))

	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{apis.URLStudent, apis.URLStudents, apis.URLStudent_id, apis.URLStudent_subjects, apis.URLGet_excel, apis.URLCreate_excel, apis.URLTeacher, apis.URLTeacher_id, apis.URLTeachers},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
			// return origin == apis.NextBaseURL
		},
		MaxAge: 12 * time.Hour,
	}))

	apis.NewStudentApi(r, db).RegisterApi()
	apis.NewTeacherApi(r, db).RegisterApi()
	apis.NewCalendarApi(r, db).RegisterApi()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, models.MsgPayload("route not implemented"))
	})

	r.Run(":8080")
}
