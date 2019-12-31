package service

import (
	"github.com/Surafeljava/Court-Case-Management-System/entity"
	notificationuse "github.com/Surafeljava/Court-Case-Management-System/notificationuse"
	//	"github.com/berrybab/entity"
)

//NotificationServiceImpl  struct
type NotificationServiceImpl struct {
	notfRepo notificationuse.NotificationRepository
}

//NewNotificationServiceImpl returns new Notification service Object
func NewNotificationServiceImpl(notf notificationuse.NotificationRepository) notificationuse.NotificationService {
	return &NotificationServiceImpl{notfRepo: notf}

}

//Notifications returns all stored Notifications from database
func (notf *NotificationServiceImpl) Notifications() ([]entity.Notification, []error) {
	notifications, errs := notf.notfRepo.Notifications()

	if len(errs) > 0 {
		return nil, errs
	}
	return notifications, errs
}

//ViewNotification retrieves a Notification  by its id(title)
func (notf *NotificationServiceImpl) ViewNotification(id string) (*entity.Notification, []error) {
	notification, errs := notf.notfRepo.ViewNotification(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return notification, errs
}

//PostNotification  admin posts a  notifiacication
func (notf *NotificationServiceImpl) PostNotification(notification *entity.Notification) (*entity.Notification, []error) {
	notif, errs := notf.notfRepo.PostNotification(notification)
	if len(errs) > 0 {
		return nil, errs
	}
	return notif, errs
}
