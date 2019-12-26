package repository

import "database/sql"

type CaseRepositoryImpl struct {
	conn *sql.DB
}
