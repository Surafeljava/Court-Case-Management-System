package models

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	NotTitle string    `gorm:"type:varchar(200);not null"`
	NotDesc  string    `gorm:"type:varchar(1500);not null"`
	NotLevel string    `gorm:"type:varchar(70);not null"`
	NotDate  time.Time `gorm:"not null"`
	//dear me!!!  remember to add &parseTime=True in the maincode
}
