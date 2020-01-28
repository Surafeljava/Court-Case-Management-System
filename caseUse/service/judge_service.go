package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type JudgeServiceImpl struct {
	judgeRepo caseUse.JudgeRepository
}

func NewJudgeServiceImpl(opRepo caseUse.JudgeRepository) *JudgeServiceImpl {
	return &JudgeServiceImpl{judgeRepo: opRepo}
}

func (osi *JudgeServiceImpl) Judges() ([]entity.Judge, error) {
	juds, errs := osi.judgeRepo.Judges()
	return juds, errs
}
func (osi *JudgeServiceImpl) Judge(id int) (*entity.Judge, []error) {
	jud, errs := osi.judgeRepo.Judge(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return jud, nil
}

func (osi *JudgeServiceImpl) CaseTypeJudges(cstype string) ([]entity.Judge, error) {
	juds, errs := osi.judgeRepo.CaseTypeJudges(cstype)
	return juds, errs
}

func (osi *JudgeServiceImpl) CreateJudge(judge *entity.Judge) (*entity.Judge, []error) {
	judge, err1 := osi.judgeRepo.CreateJudge(judge)
	if len(err1) > 0 {
		panic(err1)
	}
	return judge, err1
}

func (osi *JudgeServiceImpl) UpdateCase(judge *entity.Judge) (*entity.Judge, []error) {
	return nil, nil
}
func (osi *JudgeServiceImpl) DeleteCase(id int) error {
	return nil
}
