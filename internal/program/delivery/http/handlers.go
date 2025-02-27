package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/program"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type programHandlers struct {
	pu program.IProgramUsecase
	lg *logger.LoggerZap
}

func NewProgramHandlers(
	pu program.IProgramUsecase,
	lg *logger.LoggerZap,
) program.IProgramHandlers {
	return &programHandlers{
		pu: pu,
		lg: lg,
	}
}

func (ph *programHandlers) GetPrograms() gin.HandlerFunc {
	return func(c *gin.Context) {
		ph.lg.Info("GetPrograms called")
		programs, err := ph.pu.GetPrograms(c)
		if err != nil {
			ph.lg.Error("Error getting programs", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ph.lg.Info("GetPrograms successful")
		c.JSON(http.StatusOK, programs)
	}
}

func (ph *programHandlers) CreateProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		ph.lg.Info("CreateProgram called")
		var program models.Program
		if err := c.ShouldBindJSON(&program); err != nil {
			ph.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := ph.pu.CreateProgram(c, &program)
		if err != nil {
			ph.lg.Error("Error creating program", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ph.lg.Info("CreateProgram successful")
		c.JSON(http.StatusOK, program)
	}
}

func (ph *programHandlers) DeleteProgram() gin.HandlerFunc {
	return func(c *gin.Context) {
		ph.lg.Info("DeleteProgram called")
		programID, err := strconv.Atoi(c.Param("program_id"))
		if err != nil {
			ph.lg.Error("Invalid program ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid program ID"})
			return
		}

		err = ph.pu.DeleteProgram(c, programID)
		if err != nil {
			ph.lg.Error("Error deleting program", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ph.lg.Info("DeleteProgram successful")
		c.JSON(http.StatusOK, gin.H{"message": "Program deleted successfully"})
	}
}
