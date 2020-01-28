package service

import (
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
)

type LoginServiceImpl struct {
	loginRepo caseUse.LoginRepository
}

func NewLoginServiceImpl(logRepo caseUse.LoginRepository) *LoginServiceImpl {
	return &LoginServiceImpl{loginRepo: logRepo}
}

func (lgi *LoginServiceImpl) CheckLogin(usr *entity.UserType) (*entity.UserType, []error) {
	us, err := lgi.loginRepo.CheckLogin(usr)
	if len(err) > 0 {
		return nil, err
	}
	return us, nil
}

func (lgi *LoginServiceImpl) CheckAdmin(id string, pwd string) (*entity.Admin, []error) {
	adm, err := lgi.loginRepo.CheckAdmin(id, pwd)
	// fmt.Println(">>>>In CheckAdmin")
	if len(err) > 0 {
		return nil, err
	}
	return adm, nil
}
func (lgi *LoginServiceImpl) CheckJudge(id string, pwd string) (*entity.Judge, []error) {
	jud, err := lgi.loginRepo.CheckJudge(id, pwd)
	if len(err) > 0 {
		return nil, err
	}
	return jud, nil
}
func (lgi *LoginServiceImpl) CheckOpponent(id string, pwd string) (*entity.Opponent, []error) {
	opp, err := lgi.loginRepo.CheckOpponent(id, pwd)
	if len(err) > 0 {
		return nil, err
	}
	return opp, nil
}

func (lgi *LoginServiceImpl) GetPassword(typ int, id string) (string, error) {
	res, err := lgi.loginRepo.GetPassword(typ, id)
	return res, err
}
func (lgi *LoginServiceImpl) ChangePassword(typ int, id string, pwd string) (string, error) {
	res, err := lgi.loginRepo.ChangePassword(typ, id, pwd)
	return res, err
}
