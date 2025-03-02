package usecase

import (
	"context"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/pkg/logger"
	"github.com/opentracing/opentracing-go"
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "auditlog.GetAuditLogs")
	defer span.Finish()

	auditlogs, err := a.ar.GetAuditLogs(ctx)
	if err != nil {
		return nil, err
	}

	return auditlogs, nil
}

func (a *auditlogUsecase) GetModelAuditLogs(ctx context.Context, model string, model_id string) ([]*models.AuditLog, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auditlog.GetModelAuditLogs")
	defer span.Finish()

	auditlogs, err := a.ar.GetModelAuditLogs(ctx, model, model_id)
	if err != nil {
		return nil, err
	}

	return auditlogs, nil
}
