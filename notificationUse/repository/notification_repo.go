package repository

import (
	"github.com/Surafeljava/Court-Case-Management-System/notificationUse"

	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/jinzhu/gorm"
)

//
//NotificationRepositoryImpl struct create connection to database and implements notification repo
type NotificationRepositoryImpl struct {
	db *gorm.DB
}

//NewNotificationRepositoryImpl creates a new objct of notifaction repository
func NewNotificationRepositoryImpl(conn *gorm.DB) notificationuse.NotificationRepository {
	return &NotificationRepositoryImpl{db: conn}
}

//Notifications shows all notification posted by admin
func (notfRepo *NotificationRepositoryImpl) Notifications()([]entity.Notification,[]error) {
	notifications:=[]entity.Notification{}
	errs:=notfRepo.db.Find(&notifications).GetErrors()
	if len(errs) > 0
	{
		return nil,errs
	}
	return notifications,errs
}

//PostNotification posts notification in the database
func (notfRepo *NotificationRepositoryImpl) PostNotification(notf *entity.Notification) (*entity.Notification,[]error){
	notification:=notf
	errs:=notfRepo.db.Create(notification).GetErrors()

	if len(errs) > 0{
		return nil,errs
	}
	return notification,errs
	}

//ViewNotification retrieves a notification by title
func (notfRepo *NotificationRepositoryImpl) ViewNotification(id string) (*entity.Notification,[]error){
	notification:=entity.Notification{}
	errs:=notfRepo.db.First(&notification,id).GetErrors
	if len(errs)>0{
		return nil,errs

	}
	return &notification,errs

}
	