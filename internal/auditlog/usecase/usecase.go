package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
)

type auditlogUsecase struct {
	ar auditlog.IAuditLogRepository
	lg *logger.LoggerZap
}

func NewAuditLogUsecase(
	ar auditlog.IAuditLogRepository,
	lg *logger.LoggerZap,
) auditlog.IAuditLogUsecase {
	return &auditlogUsecase{
		ar: ar,
		lg: lg,
	}
}

func (a *auditlogUsecase) GetAuditLogs(ctx context.Context) ([]*models.AuditLog, error) {
	return a.ar.GetAuditLogs(ctx)
}

func (a *auditlogUsecase) GetModelAuditLogs(ctx context.Context, model string, model_id string) ([]*models.AuditLog, error) {
	return a.ar.GetModelAuditLogs(ctx, model, model_id)
}
