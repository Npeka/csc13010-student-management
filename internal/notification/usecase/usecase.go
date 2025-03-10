package usecase

import (
	"github.com/csc13010-student-management/internal/notification"
)

type notificationUsecase struct {
	fr notification.INotificationRepository
}

func NewNotificationUsecase(fr notification.INotificationRepository) notification.INotificationUsecase {
	return &notificationUsecase{fr: fr}
}
