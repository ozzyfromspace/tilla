package main

import (
	"log"
	"net/http"
	"os"
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

	r.Static("api/excel_files", "./excel_files")

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	apiGroup := r.Group("/api")

	apis.NewStudentApi(apiGroup, db).RegisterApi()
	apis.NewTeacherApi(apiGroup, db).RegisterApi()
	apis.NewCalendarApi(apiGroup, db).RegisterApi()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, models.MsgPayload("route not implemented"))
	})

	f, err := os.Create("sally.txt")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("Hi, I saved a plane")

	r.Run(":8080")
}
