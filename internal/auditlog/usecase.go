package auditlog

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IAuditLogUsecase interface {
	GetAuditLogs(ctx context.Context) ([]*models.AuditLog, error)
	GetModelAuditLogs(ctx context.Context, table_name string, record_id string) ([]*models.AuditLog, error)
	CreateAuditLog(ctx context.Context, auditlog *models.AuditLog) error
}
