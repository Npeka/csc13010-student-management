package auditlog

import (
	"github.com/gin-gonic/gin"
)

type IAuditLogHandlers interface {
	GetModelAuditLogs() gin.HandlerFunc
}
