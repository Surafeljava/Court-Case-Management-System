package repository

import (
	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/Surafeljava/gorm"
)

type ReportRepositoryImpl struct {
	conn *gorm.DB
}

// NewAppealGormRepo creates a new object of UserGormRepo
func NewReportGormRepo(db *gorm.DB) *ReportRepositoryImpl {
	return &ReportRepositoryImpl{conn: db}
}

func (rr *ReportRepositoryImpl) CreateCourtReport() ([]entity.Case, []entity.Judge, []error) {
	return nil, nil, nil
}
