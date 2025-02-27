package repository 
 
import "github.com/csc13010-student-management/internal/notification" 
 
type notificationRepository struct {} 
 
func NewNotificationRepository() notification.INotificationRepository { 
	return &notificationRepository{} 
} 
