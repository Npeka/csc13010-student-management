package auditlog

import (
	"github.com/gin-gonic/gin"
)

type IAuditLogHandlers interface {
	GetAuditLogs() gin.HandlerFunc
	GetModelAuditLogs() gin.HandlerFunc
}
