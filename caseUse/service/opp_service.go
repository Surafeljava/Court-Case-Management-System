package service

import (
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/Surafeljava/Court-Case-Management-System/entity"
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

func (osi *OpponentServiceImpl) CreateOpponent(opp *entity.Opponent) (*entity.Opponent, []error) {
	opp, err1 := osi.oppoRepo.CreateOpponent(opp)
	if len(err1) > 0 {
		panic(err1)
	}
	return opp, err1
}
