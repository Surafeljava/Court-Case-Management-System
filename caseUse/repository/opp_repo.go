package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/jinzhu/gorm"
)

type OpponentRepositoryImpl struct {
	conn *gorm.DB
}

func NewOpponentRepositoryImpl(Conn *gorm.DB) *OpponentRepositoryImpl {
	return &OpponentRepositoryImpl{conn: Conn}
}

func (ori *OpponentRepositoryImpl) Opponents() ([]entity.Opponent, error) {
	opp := []entity.Opponent{}
	errs := ori.conn.Find(&opp).GetErrors()
	if len(errs) > 0 {
		return nil, nil
	}
	return opp, nil
}
func (ori *OpponentRepositoryImpl) Opponent(id int) (*entity.Opponent, []error) {
	opp := entity.Opponent{}
	errs := ori.conn.First(&opp, id).GetErrors()
	if len(errs) > 0 {
		return &opp, errs
	}
	return &opp, errs
}

func (ori *OpponentRepositoryImpl) CreateOpponent(case_num string, opp *entity.Opponent) (*entity.Opponent, []error) {
	csd := opp
	errs := ori.conn.Create(&csd).GetErrors()

	if len(errs) > 0 {
		panic(errs)
		//return errs
	}

	//TODO: Add the Plaintiff to according relation column
	rel := entity.Relation{}
	err := ori.conn.Where("case_num = ?", case_num).First(&rel).GetErrors()

	if len(err) > 0 {
		return nil, err
	}

	if opp.OppType == "Plaintiff" {
		rel.PlId = opp.OppId
		ori.conn.Save(&rel)
	} else if opp.OppType == "Accused" {
		rel.AcId = opp.OppId
		ori.conn.Save(&rel)
	}

	return csd, errs
	//return nil, nil
}

func (ori *OpponentRepositoryImpl) CheckOpponentRelation(case_num string, opType string) bool {

	rel := entity.Relation{}
	err := ori.conn.Where("case_num = ?", case_num).First(&rel).GetErrors()
	if opType == "pl" {
		if rel.PlId == "notAdded" {
			return true
		} else {
			return false
		}
	} else if opType == "ac" {
		if rel.AcId == "notAdded" {
			return true
		} else {
			return false
		}
	}

	if len(err) > 0 {
		return false
	}
	return false
}
