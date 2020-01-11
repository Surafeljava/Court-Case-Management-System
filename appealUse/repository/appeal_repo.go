package repository

import "database/sql"

type AppealRepositoryImpl struct {
	conn *sql.DB
}

func NewAppealRepositoryImpl(conn *gorm.DB) appealUse.AppealRepository {
	return &AppealRepositoryImpl{db: conn}
}

func (ar *AppealRepositoryImpl) AppealGet(caseNum string) (*entity.Case, []error){
	return nil, nil
}

