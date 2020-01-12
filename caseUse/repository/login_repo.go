package repository

import (
	"reflect"

	"github.com/Surafeljava/Court-Case-Management-System/entity"
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
	errs := lgi.conn.Where("admin_id = ? and admin_pwd = ?", usr_id, usr_pwd).Find(&admin).GetErrors()
	// fmt.Println(">>>> in checkAdmin Repo")
	if len(errs) > 0 {
		return nil, errs
	}

	if reflect.DeepEqual(admin, entity.Admin{}) {
		return nil, errs
	}
	return &admin, nil
}
func (lgi *LoginRepositoryImpl) CheckJudge(id string, pwd string) (*entity.Judge, []error) {
	judge := entity.Judge{}
	usr_id := id
	usr_pwd := pwd
	errs := lgi.conn.Where("judge_id = ? and judge_pwd = ?", usr_id, usr_pwd).Find(&judge).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	if reflect.DeepEqual(judge, entity.Judge{}) {
		return nil, errs
	}
	return &judge, nil
}
func (lgi *LoginRepositoryImpl) CheckOpponent(id string, pwd string) (*entity.Opponent, []error) {
	opp := entity.Opponent{}
	usr_id := id
	usr_pwd := pwd
	errs := lgi.conn.Where("opp_id = ? and opp_pwd = ?", usr_id, usr_pwd).Find(&opp).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}

	if reflect.DeepEqual(opp, entity.Opponent{}) {
		return nil, errs
	}
	return &opp, nil
}
