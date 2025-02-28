package fileprocessor

import "github.com/gin-gonic/gin"

type IFileProcessorHandlers interface {
	ImportFile() gin.HandlerFunc
	ExportFile() gin.HandlerFunc
	// ExportStudentCertificate() gin.HandlerFunc
}
