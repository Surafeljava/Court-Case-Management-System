package repository

import "database/sql"

type ReportRepositoryImpl struct {
	conn *sql.DB
}
