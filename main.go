package main

import (
	"log"
	"net/http"
	"tilla/models"

	"github.com/gin-gonic/gin"
)

var uri = "mongodb+srv://scheduler:Schedulethejoy@eclipsecluster.mvengjn.mongodb.net/?retryWrites=true&w=majority"

func main() {
	db := &models.Database{}

	if err := db.Connect(uri); err != nil {
		log.Fatalf("[1]: %v/n", err)
	}

	log.Printf("[2]: %+v\n", db)

	r := gin.Default()
	createGroup := r.Group("create")

	createGroup.POST("/teacher", func(c *gin.Context) {
		newTeacher := models.NewTeacher()

		if err := c.BindJSON(newTeacher); err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("invalid body, could not parse teacher"))
			return
		}

		if err := db.AddTeacher(newTeacher); err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, models.MsgPayload("failed to save new teacher to database"))
			return
		}

		c.JSON(http.StatusCreated, models.MsgPayload("teacher created!"))
	})

	createGroup.POST("/student", func(c *gin.Context) {
		newStudent := models.NewStudent()

		if err := c.BindJSON(newStudent); err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("invalid body, could not parse student"))
			return
		}

		if err := db.AddStudent(newStudent); err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, models.MsgPayload("failed to save new student to database"))
			return
		}

		c.JSON(http.StatusCreated, models.MsgPayload("student created!"))
	})

	createGroup.POST("/subject", func(c *gin.Context) {
		subjectPayload := models.NewSubjectPayload()

		if err := c.BindJSON(subjectPayload); err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("invalid body, could not parse subjects"))
			return
		}

		if err := db.AddSubjects(*subjectPayload); err != nil {
			log.Println("ERR1 -", err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload(err.Error()))
			return
		}

		if len(subjectPayload.Subjects) == 1 {
			c.JSON(http.StatusCreated, models.MsgPayload("subject created!"))
			return
		}

		c.JSON(http.StatusCreated, models.MsgPayload("subjects created!"))
	})

	r.GET("/students", func(c *gin.Context) {
		students, err := db.GetStudents()

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve students"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"students": *students})
	})

	r.GET("/teachers", func(c *gin.Context) {
		teachers, err := db.GetTeachers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve teachers"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"teachers": *teachers})
	})

	r.Run(":8080")
}
