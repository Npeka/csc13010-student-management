package status

import (
	"context"

	"github.com/csc13010-student-management/internal/models"
)

type IStatusUsecase interface {
	GetStatuses(ctx context.Context) ([]*models.Status, error)
	CreateStatus(ctx context.Context, status *models.Status) error
	DeleteStatus(ctx context.Context, id int) error
}
