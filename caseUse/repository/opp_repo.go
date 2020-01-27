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

func (ori *OpponentRepositoryImpl) CreateOpponent(opp *entity.Opponent) (*entity.Opponent, []error) {
	csd := opp
	errs := ori.conn.Create(&csd).GetErrors()

	if len(errs) > 0 {
		panic(errs)
		//return errs
	}

	//TODO: Add the Plaintiff to according relation column
	rel := entity.Relation{}
	ori.conn.First(&rel, '1')

	if opp.OppType == "pl" {
		rel.PlId = opp.OppId
		ori.conn.Save(&rel)
	} else if opp.OppType == "ac" {
		rel.AcId = opp.OppId
		ori.conn.Save(&rel)
	}

	return csd, errs
	//return nil, nil
}
