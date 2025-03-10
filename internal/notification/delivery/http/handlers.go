package http

import "github.com/csc13010-student-management/internal/notification"

type notificationHandlers struct{}

func NewNotificationHandlers() notification.INotificationHandlers {
	return &notificationHandlers{}
}
