package handler

import (
	"net/http"
	"tilla/apis"
	"tilla/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// func Test(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World from Go!")
// }

// func Go(w http.ResponseWriter, r *http.Request) {
// 	app := gin.Default()
//
// 	app.GET("/api/test", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"msg": "test working..."})
// 	})
//
// 	app.ServeHTTP(w, r)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	db := models.NewDatabase()

	app := gin.New()
	gin.SetMode(gin.ReleaseMode)
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	app.Static("api/excel_files", "./excel_files")
	app.Static("api/marshmallow", "./marshmallow")

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

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "welcome to the eclipse academy API"})
	})

	apiGroup := app.Group("/api")

	apis.NewStudentApi(apiGroup, db).RegisterApi()
	apis.NewTeacherApi(apiGroup, db).RegisterApi()
	apis.NewCalendarApi(apiGroup, db).RegisterApi()

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, models.MsgPayload("route not implemented - "+c.FullPath()))
	})

	app.ServeHTTP(w, r)
}
