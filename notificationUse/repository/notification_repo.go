package repository

import "database/sql"

type NotificationRepositoryImpl struct {
	conn *sql.DB
}
