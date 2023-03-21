package handler

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
		// AllowAllOrigins: ,
		AllowOrigins:     []string{"https://eclipse-academy.vercel.app", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return true
		// },
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

	apiGroup.GET("/test", func(c *gin.Context) {
		f, err := os.Create("/tmp/sally.txt")

		if err != nil {
			log.Println("FAILED TO GET TMP FOLDER", err)
			c.JSON(http.StatusBadRequest, gin.H{"msg": "FAILED TO GET TMP FOLDER"})
		}

		n, err := f.WriteString("I saved a plane")

		if err != nil {
			log.Println("FAILED TO GET TMP FOLDER", err, n)
			c.JSON(http.StatusBadRequest, gin.H{"msg": "FAILED TO WRITE TO TMP FILE"})
		}

		c.JSON(http.StatusOK, gin.H{"msg": "done"})
	})

	apiGroup.GET("/read", func(c *gin.Context) {
		f, err := os.ReadFile("/tmp/sally.txt")

		if err != nil {
			log.Println("failed to read file from /tmp folder")
			c.JSON(http.StatusBadRequest, gin.H{"msg": "smh"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"msg": string(f)})
	})

	app.ServeHTTP(w, r)
}
