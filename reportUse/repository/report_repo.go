package repository

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/jinzhu/gorm"
)

type ReportRepositoryImpl struct {
	conn *gorm.DB
}

// NewAppealGormRepo creates a new object of UserGormRepo
func NewReportGormRepo(db *gorm.DB) *ReportRepositoryImpl {
	return &ReportRepositoryImpl{conn: db}
}

func (rr *ReportRepositoryImpl) GetClosedCases() ([]entity.Case, error) {
	cases := []entity.Case{}
	errs := rr.conn.Where("case_status = ?", "closed").Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return cases, nil
}
func (rr *ReportRepositoryImpl) GetOpenCases() ([]entity.Case, error) {
	cases := []entity.Case{}
	errs := rr.conn.Where("case_status = ?", "open").Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return cases, nil
}
func (rr *ReportRepositoryImpl) GetAllCases() ([]entity.Case, error) {
	cases := []entity.Case{}
	errs := rr.conn.Find(&cases).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return cases, nil
}
func (rr *ReportRepositoryImpl) GetCriminalJudges() ([]entity.Judge, error) {
	judges := []entity.Judge{}
	errs := rr.conn.Where("judge_type = ?", "criminal").Find(&judges).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return judges, nil
}
func (rr *ReportRepositoryImpl) GetCivilJudges() ([]entity.Judge, error) {
	judges := []entity.Judge{}
	errs := rr.conn.Where("judge_type = ?", "civil").Find(&judges).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return judges, nil
}
func (rr *ReportRepositoryImpl) GetAllJudges() ([]entity.Judge, error) {
	judges := []entity.Judge{}
	errs := rr.conn.Find(&judges).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return judges, nil
}
func (rr *ReportRepositoryImpl) GetAllNotifications() ([]entity.Notification, error) {
	notifications := []entity.Notification{}
	errs := rr.conn.Find(&notifications).GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return notifications, nil
}
