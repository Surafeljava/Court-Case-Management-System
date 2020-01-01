package repository

import (
	"database/sql"

	"github.com/Surafeljava/Court-Case-Management-System/entity"
)

type LoginRepositoryImpl struct {
	conn *sql.DB
}

func NewLoginRepositoryImpl(Conn *sql.DB) *LoginRepositoryImpl {
	return &LoginRepositoryImpl{conn: Conn}
}

func (lgi *LoginRepositoryImpl) CheckLogin(c entity.User) error {

	//TODO check the login and return something . . .

	return nil
}
