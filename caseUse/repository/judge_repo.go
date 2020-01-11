package repository

import (
	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/Surafeljava/gorm"
)

type JudgeRepositoryImpl struct {
	conn *gorm.DB
}

func NewJudgeRepositoryImpl(Conn *gorm.DB) *JudgeRepositoryImpl {
	return &JudgeRepositoryImpl{conn: Conn}
}

func (jri *JudgeRepositoryImpl) CreateJudge(judge *entity.Judge) (*entity.Judge, []error) {
	jud := judge
	errs := jri.conn.Create(&jud).GetErrors()
	if len(errs) > 0 {
		panic(errs)
		//return errs
	}
	return jud, errs
	//return nil, nil
}
