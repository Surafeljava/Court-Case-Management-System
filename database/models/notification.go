package models

import (
	"github.com/jinzhu/gorm"
)

type Notification struct {
	gorm.Model
	NotTitle string
	NotDesc  string
	NotLevel string
	NotDate  string
}
