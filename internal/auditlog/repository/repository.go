package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/auditlog"
	"github.com/csc13010-student-management/internal/models"
	"gorm.io/gorm"
)

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) auditlog.IAuditLogRepository {
	return &auditLogRepository{db: db}
}

// GetModelAuditLogs implements auditlog.IAuditLogRepository.
func (a *auditLogRepository) GetModelAuditLogs(ctx context.Context, table_name string, record_id string) ([]*models.AuditLog, error) {
	var auditLogs []*models.AuditLog
	if err := a.db.Where("table_name = ? AND record_id = ?", table_name, record_id).Find(&auditLogs).Error; err != nil {
		return nil, err
	}
	return auditLogs, nil
}
