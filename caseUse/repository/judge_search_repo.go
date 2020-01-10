package repository

import (
	"github.com/danieltefera/Project_Search/entity"
	"github.com/danieltefera/Project_Search/user"
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
func (userRepo *JudgeSearchGormRepo) Judges() ([]entity.Judge, []error) {
	cases := []entity.Judge{}
	errs := userRepo.conn.Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cases, errs
}

// Judge retrieves a user by its id from the database
func (userRepo *JudgeSearchGormRepo) Judge(id uint) (*entity.Judge, []error) {
	judge := entity.Judge{}
	errs := userRepo.conn.First(&judge, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &judge, errs
}
