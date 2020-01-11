package entity

import "github.com/Surafeljava/gorm"

type User struct {
	gorm.Model
	ID      int
	userID  string
	userPwd string
}
