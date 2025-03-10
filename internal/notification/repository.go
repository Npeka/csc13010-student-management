package notification

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type INotificationRepository interface {
	GetStatuses(ctx context.Context, statuses []uint) ([]*models.Status, error)
}
