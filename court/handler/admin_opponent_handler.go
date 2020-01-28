package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"time"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type OpponentHandler struct {
	tmpl   *template.Template
	oppSrv caseUse.OpponentService
}

func NewOpponentHandler(T *template.Template, OS caseUse.OpponentService) *OpponentHandler {
	return &OpponentHandler{tmpl: T, oppSrv: OS}
}

func (oh *OpponentHandler) NewOpponent(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		casenum := r.URL.Query().Get("case_num")
		opptype := r.URL.Query().Get("opp_type")

		if opptype == "pl" {
			opptype = "Plaintiff"
		} else if opptype == "ac" {
			opptype = "Accused"
		}

		data := struct {
			CaseNum string
			OppType string
		}{
			CaseNum: casenum,
			OppType: opptype,
		}

		oh.tmpl.ExecuteTemplate(w, "admin.newopp.layout", data)

	}

	if r.Method == http.MethodPost {

		allOpp, err1 := oh.oppSrv.Opponents()
		if err1 != nil {
			panic(err1)
		}

		csnum := r.FormValue("case_num")

		op_id := len(allOpp) + 1
		opp_id := fmt.Sprintf("OP%d", op_id)

		hasher := md5.New()
		hasher.Write([]byte("1234"))
		pwd := hex.EncodeToString(hasher.Sum(nil))

		opp_pwd := pwd

		opp_type := r.FormValue("opp_type")
		opp_name := r.FormValue("opp_fn")
		opp_gender := r.FormValue("opp_gender")
		dateFormat := "2006-01-02"
		opp_bd, _ := time.Parse(dateFormat, r.FormValue("opp_bd"))
		opp_address := r.FormValue("opp_address")
		opp_phone := r.FormValue("opp_phone")
		_, fh, _ := r.FormFile("opp_photo")

		opp_photo := fh.Filename
		newop := entity.Opponent{OppId: opp_id, OppPwd: opp_pwd, OppType: opp_type, OppName: opp_name, OppGender: opp_gender, OppBD: opp_bd, OppAddress: opp_address, OppPhone: opp_phone, OppPhoto: opp_photo}

		opp, err2 := oh.oppSrv.CreateOpponent(csnum, &newop)

		if len(err2) > 0 {
			panic(err2)
		}

		usr := entity.Messg{UserID: opp.OppId, UserName: opp.OppName, UserPwd: "1234", AddtionalMsg: "Please Change your > PASSWORD < for security purpose"}
		oh.tmpl.ExecuteTemplate(w, "admin.created.user.layout", usr)

	} else {

	}
}
