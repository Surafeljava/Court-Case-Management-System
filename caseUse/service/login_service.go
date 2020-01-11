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

func (lgi *LoginServiceImpl) CheckLogin(usr *entity.UserType) (*entity.UserType, []error) {

	//TODO check the login and return something . . .
	us, err := lgi.loginRepo.CheckLogin(usr)
	if len(err) > 0 {
		return nil, err
	}
	return us, nil
}

func (lgi *LoginServiceImpl) CheckAdmin(id string, pwd string) (*entity.Admin, []error) {
	return nil, nil
}
func (lgi *LoginServiceImpl) CheckJudge(id string, pwd string) (*entity.Judge, []error) {
	return nil, nil
}
func (lgi *LoginServiceImpl) CheckOpponent(id string, pwd string) (*entity.Opponent, []error) {
	return nil, nil
}
