package models

import (
	"github.com/jinzhu/gorm"
)

type Opponent struct {
	gorm.Model
	UserName    string
	UserId      string
	UserPwd     string
	UserGender  string
	UserBirth   string
	UserAddress string
	UserPhone   string
	UserPhoto   string
}
