package http

import (
	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/gin-gonic/gin"
)

func MapAuditLogHandlers(alGroup *gin.RouterGroup, alHandler auditlog.IAuditLogHandlers) {
	alGroup.GET("/:table_name/:record_id", alHandler.GetModelAuditLogs())
}
