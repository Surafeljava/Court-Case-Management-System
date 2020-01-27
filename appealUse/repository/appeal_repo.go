package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/appealUse"
	"github.com/jinzhu/gorm"
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

//RelationForAppeal ...
func (appealRepo *AppealGormRepo) RelationForAppeal(oppNum string) (*entity.Relation, []error) {
	rel := entity.Relation{}
	errs := appealRepo.conn.Where("pl_id =? or ac_id=?", &oppNum, &oppNum).Find(&rel).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &rel, errs
}

// CaseForAppeal retrieves a case by its id from the database
func (appealRepo *AppealGormRepo) CaseForAppeal(caseNum string) (*entity.Case, []error) {
	cas := entity.Case{}
	errs := appealRepo.conn.Where("case_num =?", &caseNum).Find(&cas).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cas, errs
}

// OppForAppeal retrieves a case by its id from the database
func (appealRepo *AppealGormRepo) OppForAppeal(oppNum string) (*entity.Opponent, []error) {
	opp := entity.Opponent{}
	errs := appealRepo.conn.Where("opp_id =?", &oppNum).Find(&opp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &opp, errs
}

// WitnessForAppeal retrieves a case by its id from the database
func (appealRepo *AppealGormRepo) WitnessForAppeal(caseNum string) (*entity.Witness, []error) {
	wit := entity.Witness{}
	errs := appealRepo.conn.Where("case_num =?", &caseNum).Find(&wit).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &wit, errs
}

// DecisionForAppeal retrieves a case by its id from the database
func (appealRepo *AppealGormRepo) DecisionForAppeal(caseNum string) (*entity.Decision, []error) {
	dic := entity.Decision{}
	errs := appealRepo.conn.Where("case_num =?", &caseNum).Find(&dic).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &dic, errs
}
