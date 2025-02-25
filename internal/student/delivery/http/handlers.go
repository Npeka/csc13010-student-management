package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/student"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type studentHandlers struct {
	su student.IStudentUsecase
	lg *logger.LoggerZap
}

func NewStudentHandlers(
	su student.IStudentUsecase,
	lg *logger.LoggerZap,
) student.IStudentHandlers {
	return &studentHandlers{
		su: su,
		lg: lg,
	}
}

// GetStudents implements student.IStudentHandlers.
func (s *studentHandlers) GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetStudents called")
		students, err := s.su.GetStudents(c)
		if err != nil {
			s.lg.Error("Error getting students", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		s.lg.Info("GetStudents successful")
		c.JSON(http.StatusOK, students)
	}
}

// CreateStudent implements student.IStudentHandlers.
func (s *studentHandlers) CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("CreateStudent called")
		var student models.Student
		if err := c.ShouldBindJSON(&student); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := s.su.CreateStudent(c, &student)
		if err != nil {
			s.lg.Error("Error creating student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("CreateStudent successful")
		c.JSON(http.StatusCreated, student)
	}
}

// UpdateStudent implements student.IStudentHandlers.
func (s *studentHandlers) UpdateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("UpdateStudent called")
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			s.lg.Error("Invalid student ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := models.Student{ID: int(id)}
		if err := c.ShouldBindJSON(&student); err != nil {
			s.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = s.su.UpdateStudent(c, &student)
		if err != nil {
			s.lg.Error("Error updating student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("UpdateStudent successful")
		c.JSON(http.StatusOK, student)
	}
}

// DeleteStudent implements student.IStudentHandlers.
func (s *studentHandlers) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("DeleteStudent called")
		id := c.Param("id")

		err := s.su.DeleteStudent(c, id)
		if err != nil {
			s.lg.Error("Error deleting student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("DeleteStudent successful")
		c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
	}
}

func (s *studentHandlers) GetOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetOptions called")
		options, err := s.su.GetOptions(c)
		if err != nil {
			s.lg.Error("Error getting options", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetOptions successful")
		c.JSON(http.StatusOK, options)
	}
}
