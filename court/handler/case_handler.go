package handler

import (
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type CaseHandler struct {
	tmpl     *template.Template
	loginSrv caseUse.LoginService
}

func NewCaseHandler(T *template.Template, LS caseUse.LoginService) *LoginHandler {
	return &LoginHandler{tmpl: T, loginSrv: LS}
}

func (lh *LoginHandler) NewCase(w http.ResponseWriter, r *http.Request) {
	lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
}
