package service

import (
	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/Surafeljava/Court-Case-Management-System/reportUse"
)

type ReportServiceImpl struct {
	reportRepo reportUse.ReportRepository
}

func NewReportServiceImpl(reportRepo reportUse.ReportRepository) *ReportServiceImpl {
	return &ReportServiceImpl{reportRepo: reportRepo}
}

func (rs *ReportServiceImpl) CreateCourtReport() ([]entity.Case, []entity.Judge, []error) {
	return nil, nil, nil
}
