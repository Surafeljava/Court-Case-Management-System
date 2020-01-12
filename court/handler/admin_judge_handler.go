package handler

import (
	"html/template"
	"net/http"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type AdminJudgeHandler struct {
	tmpl  *template.Template
	juSrv caseUse.JudgeService
}

func NewAdminJudgeHandler(T *template.Template, JS caseUse.JudgeService) *AdminJudgeHandler {
	return &AdminJudgeHandler{tmpl: T, juSrv: JS}
}

func (ajh *AdminJudgeHandler) NewJudge(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		judge_id := "JU1"
		judge_pwd := "123456"
		judge_type := r.FormValue("judge_type")
		judge_name := r.FormValue("judge_name")
		judge_gender := r.FormValue("judge_gender")
		judge_address := r.FormValue("judge_address")
		judge_phone := r.FormValue("judge_phone")
		_, fh, _ := r.FormFile("judge_photo")

		judge_photo := fh.Filename
		newop := entity.Judge{JudgeId: judge_id, JudgePwd: judge_pwd, JudgeName: judge_name, JudgeGender: judge_gender, JudgeAddress: judge_address, JudgePhone: judge_phone, JudgeType: judge_type, JudgePhoto: judge_photo}

		_, err2 := ajh.juSrv.CreateJudge(&newop)

		if len(err2) > 0 {
			panic(err2)
		}

	} else {
		ajh.tmpl.ExecuteTemplate(w, "admin.newjudge.layout", nil)
	}

	http.Redirect(w, r, "/admin/judge/new", http.StatusSeeOther)
}
