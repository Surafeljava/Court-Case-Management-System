package repository

import (
	"errors"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	notificationuse "github.com/Surafeljava/Court-Case-Management-System/notificationUse"
	"github.com/Surafeljava/gorm"
)

// MockNotificationRepo implements the notificationuse.NotificationRepository interface
type MockNotificationRepo struct {
	conn *gorm.DB
}

// NewMockNotificationRepo will create a new object of MockNotificationRepo
func NewMockNotificationRepo(db *gorm.DB) notificationuse.NotificationRepository {
	return &MockNotificationRepo{conn: db}
}

// Notifications returns all fake Notifications
func (mNotRepo *MockNotificationRepo) Notifications() ([]entity.Notification, []error) {
	nots := []entity.Notification{entity.NotificationMock}
	return nots, nil
}

// PostNotification stores a given mock Notification
func (mNotRepo *MockNotificationRepo) PostNotification(notf *entity.Notification) (*entity.Notification, []error) {
	notfs := notf
	return notfs, nil
}

// ViewNotification retrieve a fake Notification with id 1
func (mNotRepo *MockNotificationRepo) ViewNotification(id uint) (*entity.Notification, []error) {
	notf := entity.NotificationMock
	if id == 1 {
		return &notf, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateNotification updates a given fake Notification
func (mNotRepo *MockNotificationRepo) UpdateNotification(notf *entity.Notification) (*entity.Notification, []error) {
	ntfs := entity.NotificationMock
	return &ntfs, nil
}

// DeleteNotification deletes a given Notification from the database
func (mNotRepo *MockNotificationRepo) DeleteNotification(id uint) (*entity.Notification, []error) {
	notf := entity.NotificationMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &notf, nil
}
