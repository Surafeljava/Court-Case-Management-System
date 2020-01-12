package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/gorm"
)

type OpponentRepositoryImpl struct {
	conn *gorm.DB
}

func NewOpponentRepositoryImpl(Conn *gorm.DB) *OpponentRepositoryImpl {
	return &OpponentRepositoryImpl{conn: Conn}
}

func (ori *OpponentRepositoryImpl) CreateOpponent(opp *entity.Opponent) (*entity.Opponent, []error) {
	csd := opp
	errs := ori.conn.Create(&csd).GetErrors()
	if len(errs) > 0 {
		panic(errs)
		//return errs
	}
	return csd, errs
	//return nil, nil
}
