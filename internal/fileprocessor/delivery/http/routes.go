package http

import (
	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/gin-gonic/gin"
)

func MapfileProcessingHandlers(fpGroup *gin.RouterGroup, fh fileprocessor.IFileProcessorHandlers) {
	fpGroup.POST("/import", fh.ImportFile())
	fpGroup.GET("/export", fh.ExportFile())
}
