package repository

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
	"github.com/csc13010-student-management/internal/notification"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(
	db *gorm.DB,
) notification.INotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) GetStatuses(ctx context.Context, statuses []uint) ([]*models.Status, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "notificationRepository.GetStatuses")
	defer span.Finish()

	var status []*models.Status
	if err := r.db.WithContext(ctx).Where("id IN ?", statuses).Find(&status).Error; err != nil {
		return nil, errors.Wrap(err, "notificationRepository.GetStatuses.db.Find")
	}

	return status, nil
}
