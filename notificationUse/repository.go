package notificationUse

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
)

// NotificationRepository specifies notification related database operations
type NotificationRepository interface {
	ViewNotification(id uint) (*entity.Notification, []error)
	PostNotification(notf *entity.Notification) (*entity.Notification, []error)
	Notifications() ([]entity.Notification, []error)
	UpdateNotification(notf *entity.Notification) (*entity.Notification, []error)
	DeleteNotification(id uint) (*entity.Notification, []error)
}
