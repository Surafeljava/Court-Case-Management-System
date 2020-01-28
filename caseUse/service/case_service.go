package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type CaseServiceImpl struct {
	caseRepo caseUse.CaseRepository
}

func NewCaseServiceImpl(csRepo caseUse.CaseRepository) *CaseServiceImpl {
	return &CaseServiceImpl{caseRepo: csRepo}
}

//get all the cases from the database and return
func (csi *CaseServiceImpl) Cases() ([]entity.Case, error) {
	cases, errs := csi.caseRepo.Cases()
	return cases, errs
}

//get a single case
func (csi *CaseServiceImpl) Case(id int) (*entity.Case, []error) {
	css, _ := csi.caseRepo.Case(id)
	// if errs != nil {
	// 	return css, errs
	// }
	return css, nil
}

//CaseByNum ...
func (csi *CaseServiceImpl) CaseByNum(case_num string) (*entity.Case, []error) {
	css, err := csi.caseRepo.CaseByNum(case_num)
	return css, err
}

//CreateCase add new case to the database
func (csi *CaseServiceImpl) CreateCase(casedoc *entity.Case) (*entity.Case, []error) {
	CS, err1 := csi.caseRepo.CreateCase(casedoc)
	if len(err1) > 0 {
		panic(err1)
	}
	return CS, err1
}

//UpdateCase >> update a case
func (csi *CaseServiceImpl) UpdateCase(casedoc *entity.Case) (*entity.Case, []error) {
	cs, errs := csi.caseRepo.UpdateCase(casedoc)
	if len(errs) > 0 {
		return nil, errs
	}
	return cs, errs
}

//CloseCase Sevice >> close existing case
func (csi *CaseServiceImpl) CloseCase(casenum string, decision *entity.Decision) []error {
	errs := csi.caseRepo.CloseCase(casenum, decision)
	if len(errs) > 0 {
		return errs
	}
	return errs
}

//ExtendCase Service >> extend the case court date of existing case
func (csi *CaseServiceImpl) ExtendCase(casedoc *entity.Case) []error {
	return nil
}

//DeleteCase Service >> delete existing case
func (csi *CaseServiceImpl) DeleteCase(id int) (*entity.Case, []error) {
	CS, err := csi.caseRepo.DeleteCase(id)
	if len(err) > 0 {
		return nil, err
	}
	return CS, nil
}

//JudgeCases ...
func (csi *CaseServiceImpl) JudgeCases(juid string) ([]entity.Case, error) {
	cases, errs := csi.caseRepo.JudgeCases(juid)
	return cases, errs
}

//CaseJudges ...
func (csi *CaseServiceImpl) CaseJudges(case_type string) ([]entity.Judge, error) {
	juds, errs := csi.caseRepo.CaseJudges(case_type)
	return juds, errs
}
