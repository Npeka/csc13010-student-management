package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/faculty"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type facultyHandlers struct {
	fu faculty.IFacultyUsecase
	lg *logger.LoggerZap
}

func NewfacultyHandlers(
	fu faculty.IFacultyUsecase,
	lg *logger.LoggerZap,
) faculty.IFacultyHandlers {
	return &facultyHandlers{
		fu: fu,
		lg: lg,
	}
}

func (fh *facultyHandlers) GetFaculties() gin.HandlerFunc {
	return func(c *gin.Context) {
		fh.lg.Info("GetFaculties called")
		faculties, err := fh.fu.GetFaculties(c)
		if err != nil {
			fh.lg.Error("Error getting faculties", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fh.lg.Info("GetFaculties successful")
		c.JSON(http.StatusOK, faculties)
	}
}

func (fh *facultyHandlers) CreateFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		fh.lg.Info("GetFaculties called")
		var faculty models.Faculty
		if err := c.ShouldBindJSON(&faculty); err != nil {
			fh.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := fh.fu.CreateFaculty(c, &faculty)
		if err != nil {
			fh.lg.Error("Error creating faculty", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		fh.lg.Info("CreateFaculty successful")
		c.JSON(http.StatusOK, faculty)
	}
}

func (s *facultyHandlers) DeleteFaculty() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.lg.Info("DeleteFaculty called")
		faculty_id, err := strconv.Atoi(c.Param("faculty_id"))
		if err != nil {
			s.lg.Error("Invalid faculty ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid faculty ID"})
			return
		}

		err = s.fu.DeleteFaculty(c, faculty_id)
		if err != nil {
			s.lg.Error("Error deleting faculty", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		s.lg.Info("DeleteFaculty successful")
		c.JSON(http.StatusOK, gin.H{"message": "Faculty deleted successfully"})
	}
}
