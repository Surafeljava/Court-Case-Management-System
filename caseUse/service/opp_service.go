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

func (osi *OpponentServiceImpl) CreateOpponent(opp *entity.Opponent) (*entity.Opponent, []error) {
	opp, err1 := osi.oppoRepo.CreateOpponent(opp)
	if len(err1) > 0 {
		panic(err1)
	}
	return opp, err1
}
