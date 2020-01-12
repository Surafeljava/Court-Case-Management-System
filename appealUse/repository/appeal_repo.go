package repository

import (
<<<<<<< HEAD
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/appealUse"
	"github.com/Surafeljava/gorm"
)

// AppealGormRepo Implements the Repoeitory interface
type AppealGormRepo struct {
	conn *gorm.DB
}

// NewAppealGormRepo creates a new object of UserGormRepo
func NewAppealGormRepo(db *gorm.DB) user.AppealRepositroy {
	return &AppealGormRepo{conn: db}
}

// Appeal return all Cases from the database
func (appealRepo *AppealGormRepo) Appeal(oppNum string) (*entity.Case, *entity.Opponent, *entity.Witness, *entity.Decision, []error) {
	rel, err := appealRepo.RelationForAppeal(oppNum)
	if len(err) > 0 {
		return nil, nil, nil, nil, err
	}

	// Case
	caseNum := rel.CaseNum
	cases, err := appealRepo.CaseForAppeal(caseNum)
	if len(err) > 0 {
		return nil, nil, nil, nil, err
	}

	// Opponnet
	opp, err := appealRepo.OppForAppeal(oppNum)
	if len(err) > 0 {
		return nil, nil, nil, nil, err
	}

	// Witness
	wit, err := appealRepo.WitnessForAppeal(caseNum)
	if len(err) > 0 {
		return nil, nil, nil, nil, err
	}

	// Dicision
	dic, err := appealRepo.DecisionForAppeal(caseNum)
	if len(err) > 0 {
		return nil, nil, nil, nil, err
	}
	return cases, opp, wit, dic, err
}

