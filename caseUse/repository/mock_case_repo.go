package repository

import (
	"errors"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/jinzhu/gorm"
)

//MockCaseRepositoryImpl ...
type MockCaseRepositoryImpl struct {
	conn *gorm.DB
}

//NewMockCaseRepositoryImpl ...
func NewMockCaseRepositoryImpl(Conn *gorm.DB) caseUse.CaseRepository {
	return &MockCaseRepositoryImpl{conn: Conn}
}

//Cases ...
func (cri *MockCaseRepositoryImpl) Cases() ([]entity.Case, error) {
	cases := []entity.Case{entity.CaseMock}
	return cases, nil
}

//Case ...
func (cri *MockCaseRepositoryImpl) Case(id int) (*entity.Case, []error) {
	Case := entity.CaseMock
	if id == 1 {
		return &Case, nil
	}
	return nil, []error{errors.New("Not found")}
}

//CaseByNum ...
func (cri *MockCaseRepositoryImpl) CaseByNum(caseNum string) (*entity.Case, []error) {
	Case := entity.CaseMock
	if caseNum == "CS1" {
		return &Case, nil
	}
	return nil, []error{errors.New("Not found")}
}

// CreateCase ...
func (cri *MockCaseRepositoryImpl) CreateCase(casedoc *entity.Case) (*entity.Case, []error) {
	csd := casedoc
	return csd, nil
}

//UpdateCase ...
func (cri *MockCaseRepositoryImpl) UpdateCase(casedoc *entity.Case) (*entity.Case, []error) {
	cs := entity.CaseMock
	return &cs, nil
}

//CloseCase ...
func (cri *MockCaseRepositoryImpl) CloseCase(casenum string, decision *entity.Decision) []error {
	return nil
}

//ExtendCase ...
func (cri *MockCaseRepositoryImpl) ExtendCase(casedoc *entity.Case) []error {
	return nil
}

//DeleteCase ...
func (cri *MockCaseRepositoryImpl) DeleteCase(id int) (*entity.Case, []error) {
	cs := entity.CaseMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cs, nil

}

//JudgeCases ...
func (cri *MockCaseRepositoryImpl) JudgeCases(juid string) ([]entity.Case, error) {
	cases := []entity.Case{entity.CaseMock}
	return cases, nil
}

//CaseJudges ...
func (cri *MockCaseRepositoryImpl) CaseJudges(case_type string) ([]entity.Judge, error) {
	juds := []entity.Judge{entity.JudgeMock}
	return juds, nil
}
