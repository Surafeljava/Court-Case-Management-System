package repository

import "database/sql"

type AppealRepositoryImpl struct {
	conn *sql.DB
}
