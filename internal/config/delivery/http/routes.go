package http

import (
	"github.com/csc13010-student-management/internal/config"
	"github.com/gin-gonic/gin"
)

func MapconfigHandlers(cfGroup *gin.RouterGroup, ch config.IConfigHandlers) {
	cfGroup.GET("/", ch.GetConfig())
	cfGroup.PUT("/", ch.UpdateConfig())
}
