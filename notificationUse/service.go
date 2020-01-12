package notificationUse

import entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

//NotificationService interface
type NotificationService interface {
	ViewNotification(id string) (*entity.Notification, []error)
	PostNotification(notf *entity.Notification) (*entity.Notification, []error)
	Notifications() ([]entity.Notification, []error)
}
