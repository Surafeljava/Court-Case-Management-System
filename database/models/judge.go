package models

import (
	"github.com/jinzhu/gorm"
)

type Judge struct {
	gorm.Model
	JudgeName  string
	JudgeId    string
	JudgePwd   string
	JudgeType  int
	JudgePhoto string
}
