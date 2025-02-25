package http

import (
	"context"
	"net/http"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/gin-gonic/gin"
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

func (h *auditLogHandlers) GetModelAuditLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		table_name := c.Param("table_name")
		record_id := c.Param("record_id")

		auditLogs, err := h.au.GetModelAuditLogs(ctx, table_name, record_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, auditLogs)
	}
}
