package handler

import (
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/Surafeljava/Court-Case-Management-System/entity"
)

type LoginHandler struct {
	tmpl     *template.Template
	loginSrv caseUse.LoginService
}

func NewLoginHandler(T *template.Template, LS caseUse.LoginService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginSrv: LS}
}

func (lh *LoginHandler) UserLoginCheck(w http.ResponseWriter, r *http.Request) {
	//lh.tmpl.ExecuteTemplate(w, "login.layout", nil)
	if r.Method == http.MethodPost {

		user_id := r.FormValue("user_id")
		user_pwd := r.FormValue("user_pwd")

		user1 := entity.UserType{UsrId: user_id, UsrPwd: user_pwd}

		us2, err := lh.loginSrv.CheckLogin(&user1)

		if err != nil {
			panic(err)
		}

		lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", us2)

	} else {
		lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (lh *LoginHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	lh.tmpl.ExecuteTemplate(w, "login.layout", nil)
}

func CheckWho(id string) int {
	check := id[0:1]
	if check == "AD" {
		return 0
	} else if check == "JU" {
		return 1
	} else if check == "OP" {
		return 2
	}
	return -1
}
