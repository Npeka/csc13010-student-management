package http

import (
	"io"
	"net/http"

	"github.com/csc13010-student-management/internal/fileprocessor"
	"github.com/csc13010-student-management/internal/fileprocessor/processor"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type fileProcessingHandlers struct {
	fu fileprocessor.IFileProcessorUsecase
	lg *logger.LoggerZap
}

func NewFileProcessingHandlers(
	fu fileprocessor.IFileProcessorUsecase,
	lg *logger.LoggerZap,
) fileprocessor.IFileProcessorHandlers {
	return &fileProcessingHandlers{
		fu: fu,
		lg: lg,
	}
}

func (fh *fileProcessingHandlers) ImportFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "fileprocessor.ImportFile")
		defer span.Finish()

		module := c.Query("module")
		format := c.Query("format")

		file, err := c.FormFile("file")
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		fileData, err := file.Open()
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}
		defer fileData.Close()

		fileBytes, err := io.ReadAll(fileData)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		err = fh.fu.ImportFile(ctx, module, format, fileBytes)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, nil)
	}
}

func (fh *fileProcessingHandlers) ExportFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "fileprocessor.ExportFile")
		defer span.Finish()

		module := c.Query("module")
		format := c.Query("format")

		data, err := fh.fu.ExportFile(ctx, module, format)
		if err != nil {
			logger.ErrResponseWithLog(c, fh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		contentType, ext := processor.GetFileContentType(format)
		filename := module + ext

		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Data(http.StatusOK, contentType, data)
	}
}
