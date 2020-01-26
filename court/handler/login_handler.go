package handler

import (
	"fmt"
	"html/template"
	"net/http"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type LoginHandler struct {
	tmpl     *template.Template
	loginSrv caseUse.LoginService
}

func NewLoginHandler(T *template.Template, LS caseUse.LoginService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginSrv: LS}
}

func (lh *LoginHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		user_id := r.FormValue("user_id")
		user_pwd := r.FormValue("user_pwd")

		// fmt.Println(user_id)
		// fmt.Println(user_pwd)

		who := CheckWho(user_id)

		error_message := entity.SuccessMessage{Status: "Error", Message: "Wrong ID or Password Try again"}
		//success_message := entity.SuccessMessage{Status: "Success", Message: "Login Success!"}

		if who == 0 {
			adm, err := lh.loginSrv.CheckAdmin(user_id, user_pwd)
			if adm != nil {

				//Creating admin session
				CreateSession("signed_user", "admin", w)

				lh.tmpl.ExecuteTemplate(w, "admin.home.layout", adm)
			} else if len(err) > 0 {
				lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
			}
		} else if who == 1 {
			jud, err := lh.loginSrv.CheckJudge(user_id, user_pwd)
			if jud != nil {

				//Creating judge session
				CreateSession("signed_user", "judge", w)

				lh.tmpl.ExecuteTemplate(w, "judge.home.layout", jud)
			} else if len(err) > 0 {
				lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
			}
		} else if who == 2 {
			opp, err := lh.loginSrv.CheckOpponent(user_id, user_pwd)
			if opp != nil {

				//Creating opponent session
				CreateSession("signed_user", "judge", w)

				lh.tmpl.ExecuteTemplate(w, "opponent.home.layout", opp)
			} else if len(err) > 0 {
				lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
			}
		} else {
			lh.tmpl.ExecuteTemplate(w, "login.layout", error_message)
		}

	} else {
		lh.tmpl.ExecuteTemplate(w, "login.layout", nil)
	}

}

func (lh *LoginHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	//Delete the cookie created on the login process
	//back to the login page
}

func CheckWho(id string) int {
	check := id[0:2]
	fmt.Println(check)
	if check == "AD" {
		return 0
	} else if check == "JU" {
		return 1
	} else if check == "OP" {
		return 2
	}
	return -1
}

func ValidateInput(id string) {
	//Validate all the inputs in here
}

func CreateSession(name string, value string, w http.ResponseWriter) {
	c := http.Cookie{
		Name:     name,
		Value:    value,
		HttpOnly: true,
		MaxAge:   3600,
	}
	http.SetCookie(w, &c)
}
