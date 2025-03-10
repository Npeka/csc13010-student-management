package http

import (
	"net/http"

	"github.com/csc13010-student-management/internal/config"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type configHandlers struct {
	cu config.IConfigUsecase
	lg *logger.LoggerZap
}

func NewConfigHandlers(
	cu config.IConfigUsecase,
	lg *logger.LoggerZap,
) config.IConfigHandlers {
	return &configHandlers{
		cu: cu,
		lg: lg,
	}
}

func (ch *configHandlers) GetConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "configHandlers.GetConfig")
		defer span.Finish()

		cf, err := ch.cu.GetConfig(ctx)
		if err != nil {
			logger.LogResponseError(c, ch.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, cf)
	}
}

func (ch *configHandlers) UpdateConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "configHandlers.UpdateConfig")
		defer span.Finish()

		var cf models.Config
		if err := c.ShouldBindJSON(&cf); err != nil {
			logger.LogResponseError(c, ch.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := ch.cu.UpdateConfig(ctx, &cf)
		if err != nil {
			logger.LogResponseError(c, ch.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, "success")
	}
}
