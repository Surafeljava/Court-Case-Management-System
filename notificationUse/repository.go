package notificationUse

import "github.com/Surafeljava/Court-Case-Management-System/entity"

// NotificationRepository specifies notification related database operations
type NotificationRepository interface {
	ViewNotification(id string) (*entity.Notification, []error)
	PostNotification(notf *entity.Notification) (*entity.Notification, []error)
	Notifications() ([]entity.Notification, []error)
}
