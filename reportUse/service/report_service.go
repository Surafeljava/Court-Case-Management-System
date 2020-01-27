package service

import (
	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/reportUse"
)

type ReportServiceImpl struct {
	reportRepo reportUse.ReportRepository
}

func NewReportServiceImpl(reportRepo reportUse.ReportRepository) *ReportServiceImpl {
	return &ReportServiceImpl{reportRepo: reportRepo}
}

func (rs *ReportServiceImpl) GetClosedCases() ([]entity.Case, error) {
	cs, err := rs.reportRepo.GetClosedCases()
	return cs, err
}
func (rs *ReportServiceImpl) GetOpenCases() ([]entity.Case, error) {
	cs, err := rs.reportRepo.GetOpenCases()
	return cs, err
}
func (rs *ReportServiceImpl) GetAllCases() ([]entity.Case, error) {
	cs, err := rs.reportRepo.GetAllCases()
	return cs, err
}
func (rs *ReportServiceImpl) GetCriminalJudges() ([]entity.Judge, error) {
	cs, err := rs.reportRepo.GetCriminalJudges()
	return cs, err
}
func (rs *ReportServiceImpl) GetCivilJudges() ([]entity.Judge, error) {
	cs, err := rs.reportRepo.GetCivilJudges()
	return cs, err
}
func (rs *ReportServiceImpl) GetAllJudges() ([]entity.Judge, error) {
	cs, err := rs.reportRepo.GetAllJudges()
	return cs, err
}
func (rs *ReportServiceImpl) GetAllNotifications() ([]entity.Notification, error) {
	cs, err := rs.reportRepo.GetAllNotifications()
	return cs, err
}
