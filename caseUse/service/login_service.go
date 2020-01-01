package service

import (
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"

	"github.com/Surafeljava/Court-Case-Management-System/entity"
)

type LoginServiceImpl struct {
	loginRepo caseUse.LoginRepository
}

func NewLoginServiceImpl(logRepo caseUse.LoginRepository) *LoginServiceImpl {
	return &LoginServiceImpl{loginRepo: logRepo}
}

func (lgi *LoginServiceImpl) CheckLogin(c entity.User) error {

	//TODO check the login and return something . . .

	return nil
}
