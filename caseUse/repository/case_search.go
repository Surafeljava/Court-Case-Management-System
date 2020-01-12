package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/jinzhu/gorm"
)

// CaseSearchGormRepo Implements the menu.UserRepository interface
type CaseSearchGormRepo struct {
	conn *gorm.DB
}

// NewCaseSearchGormRepo creates a new object of UserGormRepo
func NewCaseSearchGormRepo(db *gorm.DB) caseUse.CaseSearchRepository {
	return &CaseSearchGormRepo{conn: db}
}

// Cases return all users from the database
func (userRepo *CaseSearchGormRepo) Cases() ([]entity.Case, []error) {
	cases := []entity.Case{}
	errs := userRepo.conn.Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cases, errs
}

// Case retrieves a user by its id from the database
func (userRepo *CaseSearchGormRepo) Case(id uint) (*entity.Case, []error) {
	cas := entity.Case{}
	errs := userRepo.conn.First(&cas, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cas, errs
}
