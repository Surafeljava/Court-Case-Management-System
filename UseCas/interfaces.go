package notification

import "github.com/courtSystem/notification/entity"

// NotificationRepository specifies notification related database operations
type NotificationRepository interface {
	ViewNotification(id string) ([]entity.Notification, error)
	PostNotification(notf entity.Notification) error
	Notifications() ([]entity.Notification, error)
}
