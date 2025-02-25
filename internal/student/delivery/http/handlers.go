package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	su student.IStudentUsecase
	lg *logger.LoggerZap
}

func NewStudentHandler(
	su student.IStudentUsecase,
	lg *logger.LoggerZap,
) student.IStudentHandler {
	return &studentHandler{
		su: su,
		lg: lg,
	}
}

// GetStudents implements student.IStudentHandler.
func (s *studentHandler) GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := s.su.GetStudents(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, students)
	}
}

// CreateStudent implements student.IStudentHandler.
func (s *studentHandler) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.su.CreateStudent(c, &student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, student)
	}
}

// UpdateStudent implements student.IStudentHandler.
func (s *studentHandler) UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := models.Student{ID: int(id)}
		if err := c.ShouldBindJSON(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = s.su.UpdateStudent(c, &student)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, student)
	}
}

// DeleteStudent implements student.IStudentHandler.
func (s *studentHandler) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := s.su.DeleteStudent(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
	}
}

func (s *studentHandler) SearchStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")

		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
			return
		}

		students, err := s.su.SearchStudents(c, query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, students)
	}
}

func (s *studentHandler) GetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		options, err := s.su.GetOptions(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, options)
	}
}