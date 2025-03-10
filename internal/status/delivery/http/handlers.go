package http

import (
	"net/http"
	"strconv"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
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
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "status.GetStatuses")
		defer span.Finish()

		statuses, err := sh.su.GetStatuses(ctx)
		if err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, statuses)
	}
}

func (sh *statusHandlers) CreateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "status.CreateStatus")
		defer span.Finish()

		var status models.Status
		if err := c.ShouldBindJSON(&status); err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err := sh.su.CreateStatus(ctx, &status)
		if err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, status)
	}
}

func (sh *statusHandlers) UpdateStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "status.UpdateStatus")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		var status models.Status
		if err := c.ShouldBindJSON(&status); err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		status.ID = uint(id)
		err = sh.su.UpdateStatus(ctx, &status)
		if err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, status)
	}
}

func (sh *statusHandlers) DeleteStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "status.DeleteStatus")
		defer span.Finish()

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusBadRequest)
			return
		}

		err = sh.su.DeleteStatus(ctx, uint(id))
		if err != nil {
			logger.LogResponseError(c, sh.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, nil)
	}
}
