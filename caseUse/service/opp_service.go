package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type OpponentServiceImpl struct {
	oppoRepo caseUse.OpponentRepository
}

func NewOpponentServiceImpl(opRepo caseUse.OpponentRepository) *OpponentServiceImpl {
	return &OpponentServiceImpl{oppoRepo: opRepo}
}

func (osi *OpponentServiceImpl) Opponents() ([]entity.Opponent, error) {
	opps, errs := osi.oppoRepo.Opponents()
	return opps, errs
}

func (osi *OpponentServiceImpl) Opponent(id int) (*entity.Opponent, []error) {
	opp, errs := osi.oppoRepo.Opponent(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return opp, nil
}

func (osi *OpponentServiceImpl) CreateOpponent(case_num string, opp *entity.Opponent) (*entity.Opponent, []error) {
	opp, err1 := osi.oppoRepo.CreateOpponent(case_num, opp)
	if len(err1) > 0 {
		panic(err1)
	}
	return opp, err1
}

func (osi *OpponentServiceImpl) CheckOpponentRelation(case_num string, opType string) bool {
	return osi.oppoRepo.CheckOpponentRelation(case_num, opType)
}
