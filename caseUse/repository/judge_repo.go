package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/jinzhu/gorm"
)

type JudgeRepositoryImpl struct {
	conn *gorm.DB
}

func NewJudgeRepositoryImpl(Conn *gorm.DB) *JudgeRepositoryImpl {
	return &JudgeRepositoryImpl{conn: Conn}
}

func (jri *JudgeRepositoryImpl) Judges() ([]entity.Judge, error) {
	juds := []entity.Judge{}
	errs := jri.conn.Find(&juds).GetErrors()
	if len(errs) > 0 {
		return nil, nil
	}
	return juds, nil
}

func (jri *JudgeRepositoryImpl) CaseTypeJudges(cstype string) ([]entity.Judge, error) {
	juds := []entity.Judge{}
	// errs := jri.conn.Find(&juds).GetErrors()
	errs := jri.conn.Where("case_type = ?", cstype).Find(&juds).GetErrors()
	if len(errs) > 0 {
		return nil, nil
	}
	return juds, nil
}

func (jri *JudgeRepositoryImpl) Judge(id int) (*entity.Judge, []error) {
	jud := entity.Judge{}
	errs := jri.conn.First(&jud, id).GetErrors()
	if len(errs) > 0 {
		return &jud, errs
	}
	return &jud, errs
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

func (jri *JudgeRepositoryImpl) UpdateCase(judge *entity.Judge) (*entity.Judge, []error) {
	return nil, nil
}

func (jri *JudgeRepositoryImpl) DeleteCase(id int) error {
	return nil
}
