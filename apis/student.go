package apis

import (
	"log"
	"net/http"
	"tilla/models"

	"github.com/gin-gonic/gin"
)

type StudentApi struct {
	r  *gin.Engine
	db *models.Database
}

func NewStudentApi(r *gin.Engine, db *models.Database) *StudentApi {
	if r == nil {
		log.Fatal("gin engine cannot be nil")
	}

	if db == nil {
		log.Fatal("database cannot be nil")
	}

	return &StudentApi{r: r, db: db}
}

func (api *StudentApi) RegisterApi() {
	api.createStudent()
	api.getStudent()
	api.getStudents()
	api.addSubjects()
}

func (api *StudentApi) createStudent() {
	api.r.POST("/student", func(c *gin.Context) {
		newStudent := models.NewStudent()

		if err := c.BindJSON(newStudent); err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("invalid body, could not parse student"))
			return
		}

		if err := api.db.AddStudent(newStudent); err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, models.MsgPayload("failed to save new student to database"))
			return
		}

		c.JSON(http.StatusCreated, models.MsgPayload("student created!"))
	})
}

func (api *StudentApi) getStudent() {
	api.r.GET("/student/:id", func(c *gin.Context) {
		studentId, found := c.Params.Get("id")

		if !found {
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("no student id provided"))
			return
		}

		student, err := api.db.GetStudentById(studentId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve student"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"student": *student})
	})
}

func (api *StudentApi) getStudents() {
	api.r.GET("/students", func(c *gin.Context) {
		students, err := api.db.GetStudents()

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve students"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"students": *students})
	})
}

func (api *StudentApi) addSubjects() {
	api.r.POST("/student/subject", func(c *gin.Context) {
		subjectPayload := models.NewSubjectPayload()

		if err := c.BindJSON(subjectPayload); err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("invalid body, could not parse subjects"))
			return
		}

		if err := api.db.AddSubjects(*subjectPayload); err != nil {
			log.Println(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload(err.Error()))
			return
		}

		if len(subjectPayload.Subjects) == 1 {
			c.JSON(http.StatusCreated, models.MsgPayload("subject created!"))
			return
		}

		c.JSON(http.StatusCreated, models.MsgPayload("subjects created!"))
	})
}
