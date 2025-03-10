package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/internal/models"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) auditlog.IAuditLogRepository {
	return &auditLogRepository{db: db}
}

func (ar *auditLogRepository) GetAuditLogs(ctx context.Context) ([]*models.AuditLog, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auditlogRepository.GetAuditLogs")
	defer span.Finish()

	var auditLogs []*models.AuditLog
	if err := ar.db.WithContext(ctx).Find(&auditLogs).Error; err != nil {
		return nil, errors.Wrap(err, "auditlog.GetAuditLogs.Find")
	}

	return auditLogs, nil
}

func (ar *auditLogRepository) GetModelAuditLogs(ctx context.Context, table_name string, record_id string) ([]*models.AuditLog, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auditlogRepository.GetModelAuditLogs")
	defer span.Finish()

	var auditLogs []*models.AuditLog
	if err := ar.db.WithContext(ctx).Where("table_name = ? AND record_id = ?", table_name, record_id).Find(&auditLogs).Error; err != nil {
		return nil, errors.Wrap(err, "auditlog.GetModelAuditLogs.Find")
	}

	return auditLogs, nil
}

func (ar *auditLogRepository) CreateAuditLog(ctx context.Context, auditlog *models.AuditLog) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "auditlogRepository.CreateAuditLog")
	defer span.Finish()

	if err := ar.db.WithContext(ctx).Create(auditlog).Error; err != nil {
		return errors.Wrap(err, "auditlog.CreateAuditLog.Create")
	}

	return nil
}
