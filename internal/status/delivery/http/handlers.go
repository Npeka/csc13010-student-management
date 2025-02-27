package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type statusHandlers struct {
	su status.IStatusUsecase
	lg *logger.LoggerZap
}

func NewStatusHandlers(
	su status.IStatusUsecase,
	lg *logger.LoggerZap,
) status.IStatusHandlers {
	return &statusHandlers{
		su: su,
		lg: lg,
	}
}

func (sh *statusHandlers) GetStatuses() gin.HandlerFunc {
	return func(c *gin.Context) {
		sh.lg.Info("GetFaculties called")
		statuses, err := sh.su.GetStatuses(c)
		if err != nil {
			sh.lg.Error("Error getting statuses", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sh.lg.Info("GetStatuses successful")
		c.JSON(http.StatusOK, statuses)
	}
}

func (sh *statusHandlers) CreateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		sh.lg.Info("GetFaculties called")
		var status models.Status
		if err := c.ShouldBindJSON(&status); err != nil {
			sh.lg.Error("Error binding JSON", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := sh.su.CreateStatus(c, &status)
		if err != nil {
			sh.lg.Error("Error creating status", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sh.lg.Info("GetFaculties successful")
		c.JSON(http.StatusOK, status)
	}
}

func (sh *statusHandlers) DeleteStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		sh.lg.Info("DeleteStatus called")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			sh.lg.Error("Invalid status ID", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status ID"})
			return
		}

		err = sh.su.DeleteStatus(c, id)
		if err != nil {
			sh.lg.Error("Error deleting status", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		sh.lg.Info("DeleteStatus successful")
		c.JSON(http.StatusOK, gin.H{"message": "Status deleted successfully"})
	}
}
