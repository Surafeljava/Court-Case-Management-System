package entity

import "github.com/Surafeljava/gorm"

//Judge struct
type Judge struct {
	gorm.Model
	JudgeName  string
	JudgeId    string
	JudgePwd   string
	JudgeType  int
	JudgePhoto string
}
