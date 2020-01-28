package repository

import (
	"crypto/md5"
	"encoding/hex"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/jinzhu/gorm"
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

	if pwdnew == judge.JudgePwd {
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

	if pwdnew == opp.OppPwd {
		return &opp, nil
	}

	return nil, nil
}

func (lgi *LoginRepositoryImpl) GetPassword(typ int, id string) (string, error) {

	// uid := id
	if typ == 0 {
		admin := entity.Admin{}
		errs := lgi.conn.Where("admin_id = ?", id).Find(&admin).GetErrors()
		if len(errs) > 0 {
			return "", errs[0]
		}
		return admin.AdminPwd, nil
	} else if typ == 1 {
		judge := entity.Judge{}
		errs := lgi.conn.Where("judge_id = ?", id).Find(&judge).GetErrors()
		if len(errs) > 0 {
			return "", errs[0]
		}
		return judge.JudgePwd, nil
	} else {
		opp := entity.Opponent{}
		errs := lgi.conn.Where("opp_id = ?", id).Find(&opp).GetErrors()
		if len(errs) > 0 {
			return "", errs[0]
		}
		return opp.OppPwd, nil
	}
}
func (lgi *LoginRepositoryImpl) ChangePassword(typ int, id string, pwd string) (string, error) {
	if typ == 0 {
		admin := entity.Admin{}
		errs := lgi.conn.Where("admin_id = ?", id).Find(&admin).GetErrors()
		if len(errs) > 0 {
			return "error changing password", errs[0]
		}

		admin.AdminPwd = HashPwd(pwd)
		err := lgi.conn.Where("admin_id = ?", id).Save(&admin).GetErrors()
		if len(err) > 0 {
			return "error changing password", errs[0]
		}

		return "Password Changed!", nil
	} else if typ == 1 {
		judge := entity.Judge{}
		errs := lgi.conn.Where("judge_id = ?", id).Find(&judge).GetErrors()
		if len(errs) > 0 {
			return "error changing password", errs[0]
		}

		judge.JudgePwd = HashPwd(pwd)
		err := lgi.conn.Where("judge_id = ?", id).Save(&judge).GetErrors()
		if len(err) > 0 {
			return "error changing password", errs[0]
		}

		return "Password Changed!", nil
	} else {
		opp := entity.Opponent{}
		errs := lgi.conn.Where("opp_id = ?", id).Find(&opp).GetErrors()
		if len(errs) > 0 {
			return "error changing password", errs[0]
		}

		opp.OppPwd = HashPwd(pwd)
		err := lgi.conn.Where("opp_id = ?", id).Save(&opp).GetErrors()
		if len(err) > 0 {
			return "error changing password", errs[0]
		}

		return "Password Changed!", nil
	}
}

func HashPwd(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdnew := hex.EncodeToString(hasher.Sum(nil))

	return pwdnew
}
