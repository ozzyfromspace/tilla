package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func Test(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World from Go!")
// }

func Go(w http.ResponseWriter, r *http.Request) {
	app := gin.Default()

	app.GET("/api/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "test working..."})
	})

	app.ServeHTTP(w, r)
}

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	db := models.NewDatabase()
//
// 	app := gin.Default()
// 	app.Use(gzip.Gzip(gzip.DefaultCompression))
//
// 	app.Use(cors.New(cors.Config{
// 		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
// 		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
// 		AllowCredentials: true,
// 		AllowOriginFunc: func(origin string) bool {
// 			return true
// 		},
// 		MaxAge: 12 * time.Hour,
// 	}))
//
// 	app.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"msg": "welcome to the eclipse academy API"})
// 	})
//
// 	apiGroup := app.Group("/api")
//
// 	apis.NewStudentApi(apiGroup, db).RegisterApi()
// 	apis.NewTeacherApi(apiGroup, db).RegisterApi()
// 	apis.NewCalendarApi(apiGroup, db).RegisterApi()
//
// 	app.NoRoute(func(c *gin.Context) {
// 		payload := gin.H{
// 			"msg":  "route not implemented",
// 			"path": c.Request.URL.Path,
// 		}
// 		// c.JSON(http.StatusNotImplemented, models.MsgPayload("route not implemented"))
// 		c.JSON(http.StatusNotImplemented, payload)
// 	})
//
// 	// apiGroup.Run(":8080")
// 	app.ServeHTTP(w, r)
// }
//
