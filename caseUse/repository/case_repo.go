package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/jinzhu/gorm"
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

func (cri *CaseRepositoryImpl) CaseByNum(case_num string) (*entity.Case, []error) {
	css := entity.Case{}
	errs := cri.conn.Where("case_num = ?", case_num).First(&css).GetErrors()
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

	relation := entity.Relation{CaseNum: csd.CaseNum, PlId: "", AcId: ""}
	cri.conn.Create(&relation)

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
func (cri *CaseRepositoryImpl) CloseCase(casenum string, decision *entity.Decision) []error {
	cs := entity.Case{}
	//errs := cri.conn.Model(&cs).Where("case_num = ?", cs.CaseNum).Update("case_status", "Closed").GetErrors()
	errs := cri.conn.Where("case_num = ?", casenum).First(&cs)
	cs.CaseStatus = "Closed"
	er := cri.conn.Save(&cs).GetErrors()

	if len(er) > 0 {
		return er
	}

	errs2 := cri.conn.Save(&decision).GetErrors()

	if errs != nil || len(errs2) > 0 {
		return errs2
	}
	return nil
}
func (cri *CaseRepositoryImpl) ExtendCase(casedoc *entity.Case) []error {
	cs := casedoc
	errs := cri.conn.Save(&cs).GetErrors()
	if len(errs) > 0 {
		return errs
	}
	return errs
}
func (cri *CaseRepositoryImpl) DeleteCase(id int) []error {
	cs, err := cri.Case(id)
	if len(err) > 0 {
		return err
	}
	errs := cri.conn.Delete(cs, id).GetErrors()
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (cri *CaseRepositoryImpl) JudgeCases(juid string) ([]entity.Case, error) {
	cases := []entity.Case{}
	errs := cri.conn.Model(&cases).Where("case_judge = ? AND case_status = ?", juid, "open").Find(&cases).GetErrors()
	if len(errs) > 0 {

		return nil, nil
	}
	return cases, nil
}

func (cri *CaseRepositoryImpl) CaseJudges(case_type string) ([]entity.Judge, error) {
	juds := []entity.Judge{}
	errs := cri.conn.Find(&juds).GetErrors()
	if len(errs) > 0 {
		return nil, nil
	}
	return juds, nil
}
