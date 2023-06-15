package api

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseAPI interface {
	AddCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
}

type courseAPI struct {
	courseRepo repo.CourseRepository
}

func NewCourseAPI(courseRepo repo.CourseRepository) *courseAPI {
	return &courseAPI{courseRepo}
}

func (cr *courseAPI) AddCourse(c *gin.Context) {
	var newCourse model.Course
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := cr.courseRepo.Store(&newCourse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add course success"})
}

func (cr *courseAPI) DeleteCourse(c *gin.Context) {
	courseID := c.Param("id")

	// Validate the courseID parameter
	id, err := strconv.Atoi(courseID)
	if err != nil {
		// Return a JSON response with a 400 Bad Request status code and an error message
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	// Delete the course from the database using the repository
	err = cr.courseRepo.Delete(id)
	if err != nil {
		// Return a JSON response with a 500 Internal Server Error status code and an error message
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a JSON response with a 200 OK status code and a success message
	c.JSON(http.StatusOK, gin.H{"message": "Course delete success"})
}
