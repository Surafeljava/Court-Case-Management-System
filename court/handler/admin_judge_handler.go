package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

		allJudges, err1 := ajh.juSrv.Judges()
		if err1 != nil {
			panic(err1)
		}

		ju_id := len(allJudges) + 1
		judge_id := fmt.Sprintf("JU%d", ju_id)

		hasher := md5.New()
		hasher.Write([]byte("1234"))
		pwd := hex.EncodeToString(hasher.Sum(nil))

		judge_pwd := pwd

		judge_type := r.FormValue("judge_type")
		judge_name := r.FormValue("judge_name")
		judge_gender := r.FormValue("judge_gender")
		judge_address := r.FormValue("judge_address")
		judge_phone := r.FormValue("judge_phone")
		_, fh, _ := r.FormFile("judge_photo")

		judge_photo := fh.Filename
		newjd := entity.Judge{JudgeId: judge_id, JudgePwd: judge_pwd, JudgeName: judge_name, JudgeGender: judge_gender, JudgeAddress: judge_address, JudgePhone: judge_phone, JudgeType: judge_type, JudgePhoto: judge_photo}

		jd, err2 := ajh.juSrv.CreateJudge(&newjd)

		if len(err2) > 0 {
			panic(err2)
		}

		usr := entity.Messg{UserID: jd.JudgeId, UserName: jd.JudgeName, UserPwd: "1234", AddtionalMsg: "Please Change your > PASSWORD < for security purpose"}
		ajh.tmpl.ExecuteTemplate(w, "admin.created.user.layout", usr)

	} else {
		ajh.tmpl.ExecuteTemplate(w, "admin.newjudge.layout", nil)
	}

}
