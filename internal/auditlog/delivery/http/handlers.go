package http

import (
	"net/http"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/csc13010-student-management/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type auditLogHandlers struct {
	au auditlog.IAuditLogUsecase
	lg *logger.LoggerZap
}

func NewAuditLogHandlers(
	au auditlog.IAuditLogUsecase,
	lg *logger.LoggerZap,
) auditlog.IAuditLogHandlers {
	return &auditLogHandlers{
		au: au,
		lg: lg,
	}
}

func (h *auditLogHandlers) GetAuditLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "auditlog.GetAuditLogs")
		defer span.Finish()

		auditLogs, err := h.au.GetAuditLogs(ctx)
		if err != nil {
			logger.ErrResponseWithLog(c, h.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, auditLogs)
	}
}

func (h *auditLogHandlers) GetModelAuditLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "auditlog.GetModelAuditLogs")
		defer span.Finish()

		table_name := c.Param("table_name")
		record_id := c.Param("record_id")

		auditLogs, err := h.au.GetModelAuditLogs(ctx, table_name, record_id)
		if err != nil {
			logger.ErrResponseWithLog(c, h.lg, err)
			response.Error(c, http.StatusInternalServerError)
			return
		}

		response.Success(c, response.ErrCodeSuccess, auditLogs)
	}
}
