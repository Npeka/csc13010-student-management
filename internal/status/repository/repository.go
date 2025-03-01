package repository

import (
	"context"
	"fmt"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"gorm.io/gorm"
)

type statusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(
	db *gorm.DB,
) status.IStatusRepository {
	return &statusRepository{
		db: db,
	}
}

func (s *statusRepository) GetStatuses(ctx context.Context) ([]*models.Status, error) {
	var statuses []*models.Status
	if err := s.db.WithContext(ctx).Find(&statuses).Error; err != nil {
		return nil, fmt.Errorf("failed to get statuses: %w", err)
	}
	return statuses, nil
}

func (s *statusRepository) CreateStatus(ctx context.Context, status *models.Status) error {
	if err := s.db.WithContext(ctx).Create(status).Error; err != nil {
		return fmt.Errorf("failed to create status: %w", err)
	}
	return nil
}

func (s *statusRepository) DeleteStatus(ctx context.Context, id uint) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Status{}).Error; err != nil {
		return fmt.Errorf("failed to delete status: %w", err)
	}
	return nil
}
