package apis

import (
	"log"
	"net/http"
	"tilla/models"

	"github.com/gin-gonic/gin"
)

type TeacherApi struct {
	r  *gin.RouterGroup
	db *models.Database
}

func NewTeacherApi(r *gin.RouterGroup, db *models.Database) *TeacherApi {
	if r == nil {
		log.Fatal("gin engine cannot be nil")
	}

	if db == nil {
		log.Fatal("database cannot be nil")
	}

	return &TeacherApi{r: r, db: db}
}

func (api *TeacherApi) RegisterApi() {
	api.getTeacher()
	api.getTeachers()
	api.createTeacher()
}

func (api *TeacherApi) createTeacher() {
	api.r.POST(RTeacher, func(c *gin.Context) {
		newTeacher := models.NewTeacher()

		if err := c.BindJSON(newTeacher); err != nil {
			log.Print(err)
			c.JSON(http.StatusUnprocessableEntity, models.MsgPayload("invalid body, could not parse teacher"))
			return
		}

		if err := api.db.AddTeacher(newTeacher); err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, models.MsgPayload("failed to save new teacher to database"))
			return
		}

		c.JSON(http.StatusCreated, models.MsgPayload("teacher created!"))
	})
}

func (api *TeacherApi) getTeachers() {
	api.r.GET(RTeachers, func(c *gin.Context) {
		teachers, err := api.db.GetTeachers()

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve teachers"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"teachers": *teachers})
	})
}

func (api *TeacherApi) getTeacher() {
	api.r.GET(RTeacher_id, func(c *gin.Context) {
		studentId, found := c.Params.Get("id")

		if !found {
			c.JSON(http.StatusBadRequest, models.MsgPayload("no teacher id provided"))
			return
		}

		student, err := api.db.GetTeacherById(studentId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MsgPayload("could not retrieve teacher"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"teacher": *student})
	})
}
