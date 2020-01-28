package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	notificationuse "github.com/Surafeljava/Court-Case-Management-System/notificationUse"
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
func (notf *NotificationServiceImpl) ViewNotification(id uint) (*entity.Notification, []error) {
	notification, errs := notf.notfRepo.ViewNotification(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return notification, errs
}

//PostNotification  admin posts a  notifiacication
func (notf *NotificationServiceImpl) PostNotification(notification *entity.Notification) (*entity.Notification, []error) {
	ntf, errs := notf.notfRepo.PostNotification(notification)
	if len(errs) > 0 {
		return nil, errs
	}
	return ntf, errs
}

//UpdateNotification implemented below
func (notf *NotificationServiceImpl) UpdateNotification(notification *entity.Notification) (*entity.Notification, []error) {
	notfic, errs := notf.notfRepo.UpdateNotification(notification)

	if len(errs) > 0 {
		return nil, errs
	}
	return notfic, errs
}

//DeleteNotification deletes a given notification
func (notf *NotificationServiceImpl) DeleteNotification(id uint) (*entity.Notification, []error) {
	ntf, errs := notf.notfRepo.DeleteNotification(id)
	if len(errs) > 0 {
		return nil, errs

	}
	return ntf, nil
}
