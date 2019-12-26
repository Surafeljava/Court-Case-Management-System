package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	CaseNum   string
	CaseTitle string
	CaseDesc  string
	CaseStat  string
	CaseType  string
	CaseJudge string
}
