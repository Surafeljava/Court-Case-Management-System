package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/gorm"
)

type CaseRepositoryImpl struct {
	conn *gorm.DB
}

func NewCaseRepositoryImpl(Conn *gorm.DB) *CaseRepositoryImpl {
	return &CaseRepositoryImpl{conn: Conn}
}

func (cri *CaseRepositoryImpl) Cases() ([]entity.Case, error) {
	cases := []entity.Case{}
	errs := cri.conn.Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, nil
	}
	return cases, nil
}
func (cri *CaseRepositoryImpl) Case(id int) (*entity.Case, []error) {
	css := entity.Case{}
	errs := cri.conn.First(&css, id).GetErrors()
	if len(errs) > 0 {
		return &css, errs
	}
	return &css, errs
}
func (cri *CaseRepositoryImpl) CreateCase(casedoc *entity.Case) []error {
	csd := casedoc
	errs := cri.conn.Create(&csd).GetErrors()
	if len(errs) > 0 {
		panic(errs)
		//return errs
	}
	return errs
}
func (cri *CaseRepositoryImpl) UpdateCase(casedoc *entity.Case) (*entity.Case, []error) {
	cs := casedoc
	errs := cri.conn.Save(&cs).GetErrors()
	//errs := cri.conn.Model(&cs).Where("case_num = ?", cs.CaseNum).Update("case_title", "case_desc").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cs, errs
}
func (cri *CaseRepositoryImpl) CloseCase(casedoc entity.Case) error {
	return nil
}
func (cri *CaseRepositoryImpl) ExtendCase(casedoc entity.Case) error {
	return nil
}
func (cri *CaseRepositoryImpl) DeleteCase(id int) error {
	return nil
}
