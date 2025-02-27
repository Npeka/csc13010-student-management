package http

import (
	"github.com/csc13010-student-management/internal/status"
	"github.com/gin-gonic/gin"
)

func MapstatusHandlers(ftGroup *gin.RouterGroup, h status.IStatusHandlers) {
	ftGroup.GET("/", h.GetStatuses())
	ftGroup.POST("/", h.CreateStatus())
	ftGroup.DELETE("/:id", h.DeleteStatus())
}
