package repository

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/gorm"
)

type LoginRepositoryImpl struct {
	conn *gorm.DB
}

func NewLoginRepositoryImpl(Conn *gorm.DB) *LoginRepositoryImpl {
	return &LoginRepositoryImpl{conn: Conn}
}

func (lgi *LoginRepositoryImpl) CheckLogin(usr *entity.UserType) (*entity.UserType, []error) {

	//TODO check the login and return something . . .
	user2 := entity.UserType{}
	usr_id := usr.UsrId
	usr_pwd := usr.UsrPwd
	//errs := lgi.conn.Find(&user2, usr_id, usr_pwd).GetErrors()
	errs := lgi.conn.Where("usr_id = ?", usr_id).Find(&user2).GetErrors()
	//errs := lgi.conn.Find(&user2).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	if user2.UsrPwd == usr_pwd {
		return &user2, errs
	}
	return &user2, errs

}

func (lgi *LoginRepositoryImpl) CheckAdmin(id string, pwd string) (*entity.Admin, []error) {
	admin := entity.Admin{}
	usr_id := id
	usr_pwd := pwd
	errs := lgi.conn.Where("admin_id = ?", usr_id).Find(&admin).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	hasher := md5.New()
	hasher.Write([]byte(usr_pwd))
	pwdnew := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println(pwdnew)
	fmt.Println(admin.AdminPwd)

	if pwdnew == admin.AdminPwd {
		return &admin, nil
	}

	return nil, nil
}
func (lgi *LoginRepositoryImpl) CheckJudge(id string, pwd string) (*entity.Judge, []error) {
	judge := entity.Judge{}
	usr_id := id
	usr_pwd := pwd
	errs := lgi.conn.Where("judge_id = ?", usr_id).Find(&judge).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	hasher := md5.New()
	hasher.Write([]byte(usr_pwd))
	pwdnew := hex.EncodeToString(hasher.Sum(nil))

	if pwdnew == judge.JudgeId {
		return &judge, nil
	}

	return nil, nil
}
func (lgi *LoginRepositoryImpl) CheckOpponent(id string, pwd string) (*entity.Opponent, []error) {
	opp := entity.Opponent{}
	usr_id := id
	usr_pwd := pwd
	errs := lgi.conn.Where("opp_id = ?", usr_id).Find(&opp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	hasher := md5.New()
	hasher.Write([]byte(usr_pwd))
	pwdnew := hex.EncodeToString(hasher.Sum(nil))

	if pwdnew == opp.OppId {
		return &opp, nil
	}

	return nil, nil
}
