package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
	"github.com/jinzhu/gorm"
)

// JudgeSearchGormRepo Implements the menu.UserRepository interface
type JudgeSearchGormRepo struct {
	conn *gorm.DB
}

// NewJudgeSearchGormRepo creates a new object of UserGormRepo
func NewJudgeSearchGormRepo(db *gorm.DB) user.JudgeSearchRepository {
	return &JudgeSearchGormRepo{conn: db}
}

// Judges return all users from the database
func (judgeRepo *JudgeSearchGormRepo) Judges() ([]entity.Judge, []error) {
	cases := []entity.Judge{}
	errs := judgeRepo.conn.Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cases, errs
}

// Judge retrieves a user by its id from the database
func (judgeRepo *JudgeSearchGormRepo) Judge(judgeID string) (*entity.Judge, []error) {
	judge := entity.Judge{}
	errs := judgeRepo.conn.Where("judge_id =?", &judgeID).Find(&judge).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &judge, errs
}
