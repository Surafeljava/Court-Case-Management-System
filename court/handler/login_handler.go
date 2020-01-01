package handler

import (
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type LoginHandler struct {
	tmpl     *template.Template
	loginSrv caseUse.LoginService
}

func NewLoginHandler(T *template.Template, LS caseUse.LoginService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginSrv: LS}
}

func (lh *LoginHandler) UserLoginCheck(w http.ResponseWriter, r *http.Request) {
	lh.tmpl.ExecuteTemplate(w, "login.layout", nil)
}
