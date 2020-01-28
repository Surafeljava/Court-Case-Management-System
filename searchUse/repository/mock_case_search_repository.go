package repository

import (
	"errors"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
	"github.com/Surafeljava/gorm"
)

// MockCaseSearchRepo implements the menu.CategoryRepository interface
type MockCaseSearchRepo struct {
	conn *gorm.DB
}

// NewMockCaseSearchRepo will create a new object of MockCaseSearchRepo
func NewMockCaseSearchRepo(db *gorm.DB) user.CaseSearchRepository {
	return &MockCaseSearchRepo{conn: db}
}

// Cases returns all fake Cases
func (mCasRepo *MockCaseSearchRepo) Cases() ([]entity.Case, []error) {
	cases := []entity.Case{entity.CaseMock}
	return cases, nil
}

// Case retrieves a fake Case with caseNum "CS1"
func (mCasRepo *MockCaseSearchRepo) Case(caseNum string) (*entity.Case, []error) {
	Case := entity.CaseMock
	if caseNum == "CS1" {
		return &Case, nil
	}
	return nil, []error{errors.New("Not found")}
}
