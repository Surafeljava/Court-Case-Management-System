package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
	"github.com/jinzhu/gorm"
)

// CaseSearchGormRepo Implements the Repository interface
type CaseSearchGormRepo struct {
	conn *gorm.DB
}

// NewCaseSearchGormRepo creates a new object of CaseSearchGormRepo
func NewCaseSearchGormRepo(db *gorm.DB) user.CaseSearchRepository {
	return &CaseSearchGormRepo{conn: db}
}

// Cases return all Cases from the database
func (caseRepo *CaseSearchGormRepo) Cases() ([]entity.Case, []error) {
	cases := []entity.Case{}
	errs := caseRepo.conn.Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cases, errs
}

// Case retrieves a case by its id from the database
func (caseRepo *CaseSearchGormRepo) Case(caseNum string) (*entity.Case, []error) {
	cas := entity.Case{}
	errs := caseRepo.conn.Where("case_num =?", &caseNum).Find(&cas).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cas, errs
}
