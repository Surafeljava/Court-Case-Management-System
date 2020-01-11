package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
	"github.com/Surafeljava/Court-Case-Management-System/entity"
)

type OpponentHandler struct {
	tmpl   *template.Template
	oppSrv caseUse.OpponentService
}

func NewOpponentHandler(T *template.Template, OS caseUse.OpponentService) *OpponentHandler {
	return &OpponentHandler{tmpl: T, oppSrv: OS}
}

func (oh *OpponentHandler) NewOpponent(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		opp_id := "OP1"
		opp_pwd := "123456"
		opp_type := r.FormValue("opp_type")
		opp_name := r.FormValue("opp_fn")
		opp_gender := r.FormValue("opp_gender")
		layoutISO := "2006-01-02"
		opp_bd, _ := time.Parse(layoutISO, r.FormValue("opp_bd"))
		//opp_bd := r.FormValue("opp_bd")
		fmt.Println(r.FormValue("opp_bd"))
		opp_address := r.FormValue("opp_address")
		opp_phone := r.FormValue("opp_phone")
		_, fh, _ := r.FormFile("opp_photo")

		opp_photo := fh.Filename
		newop := entity.Opponent{OppId: opp_id, OppPwd: opp_pwd, OppType: opp_type, OppName: opp_name, OppGender: opp_gender, OppBD: opp_bd, OppAddress: opp_address, OppPhone: opp_phone, OppPhoto: opp_photo}

		_, err2 := oh.oppSrv.CreateOpponent(&newop)

		if len(err2) > 0 {
			panic(err2)
		}

	} else {
		oh.tmpl.ExecuteTemplate(w, "admin.newopp.layout", nil)
	}

	http.Redirect(w, r, "/admin/opponent/new", http.StatusSeeOther)
}
