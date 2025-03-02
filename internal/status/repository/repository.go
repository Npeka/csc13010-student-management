package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/status"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepository.GetStatuses")
	defer span.Finish()

	var statuses []*models.Status
	if err := s.db.WithContext(ctx).Find(&statuses).Error; err != nil {
		return nil, errors.Wrap(err, "statusRepository.GetStatuses.Find")
	}
	return statuses, nil
}

func (s *statusRepository) CreateStatus(ctx context.Context, status *models.Status) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepository.CreateStatus")
	defer span.Finish()

	if err := s.db.WithContext(ctx).Create(status).Error; err != nil {
		return errors.Wrap(err, "statusRepository.CreateStatus.Create")
	}
	return nil
}

func (s *statusRepository) UpdateStatus(ctx context.Context, status *models.Status) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepository.UpdateStatus")
	defer span.Finish()

	if err := s.db.WithContext(ctx).Save(status).Error; err != nil {
		return errors.Wrap(err, "statusRepository.UpdateStatus.Save")
	}
	return nil
}

func (s *statusRepository) DeleteStatus(ctx context.Context, id uint) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "statusRepository.DeleteStatus")
	defer span.Finish()

	if err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Status{}).Error; err != nil {
		return errors.Wrap(err, "statusRepository.DeleteStatus.Delete")
	}
	return nil
}
