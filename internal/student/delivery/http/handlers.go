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

// GetStudentByStudentID implements student.IStudentHandlers.
func (s *studentHandlers) GetStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetStudentByStudentID called")
		student_id := c.Param("student_id")

		student, err := s.su.GetStudentByStudentID(c, student_id)
		if err != nil {
			s.lg.Error("Error getting student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetStudentByStudentID successful")
		c.JSON(http.StatusOK, student)
	}
}

// GetFullInfoStudentByStudentID implements student.IStudentHandlers.
func (s *studentHandlers) GetFullInfoStudentByStudentID() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("GetFullInfoStudentByStudentID called")
		student_id := c.Param("student_id")

		student, err := s.su.GetFullInfoStudentByStudentID(c, student_id)
		if err != nil {
			s.lg.Error("Error getting full info student", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("GetFullInfoStudentByStudentID successful")
		c.JSON(http.StatusOK, student)
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
		student_id, err := strconv.ParseInt(c.Param("student_id"), 10, 64)
		if err != nil {
			s.lg.Error("Invalid student ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		student := models.Student{StudentID: strconv.FormatInt(student_id, 10)}
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

// const defaultDeleteTimeLimit = 30 * time.Minute

// DeleteStudent implements student.IStudentHandlers.
func (s *studentHandlers) DeleteStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if os.Getenv("ALLOW_DELETE_ANYTIME") == "false" {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "DeleteStudent is disabled"})
		// 	return
		// }

		s.lg.Info("DeleteStudent called")
		student_id := c.Param("student_id")

		// student, err := s.su.GetStudentByStudentID(c, student_id)
		// if err != nil {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		// 	return
		// }

		// deleteTimeLimit := defaultDeleteTimeLimit
		// if envLimit, exists := os.LookupEnv("DELETE_TIME_LIMIT"); exists {
		// 	parsedLimit, err := time.ParseDuration(envLimit)
		// 	if err == nil {
		// 		deleteTimeLimit = parsedLimit
		// 	}
		// }

		// if time.Since(student.CreatedAt) > deleteTimeLimit {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Student cannot be deleted after the allowed time"})
		// 	return
		// }

		err := s.su.DeleteStudent(c, student_id)
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
