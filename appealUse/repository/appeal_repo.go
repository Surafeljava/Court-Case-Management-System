package repository

import (
	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/Surafeljava/gorm"
)

type AppealRepositoryImpl struct {
	conn *gorm.DB
}

func NewAppealRepositoryImpl(conn *gorm.DB) *AppealRepositoryImpl {
	return &AppealRepositoryImpl{conn: conn}
}

func (ar *AppealRepositoryImpl) AppealGet(caseNum string) (*entity.Case, []error) {
	return nil, nil
}
