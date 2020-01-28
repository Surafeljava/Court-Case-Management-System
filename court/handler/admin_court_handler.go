package handler

import (
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"net/http"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type AdminCourtHandler struct {
	tmpl      *template.Template
	courtServ caseUse.CourtService
}

func NewAdminCourtHandler(T *template.Template, JS caseUse.CourtService) *AdminCourtHandler {
	return &AdminCourtHandler{tmpl: T, courtServ: JS}
}

func (ach *AdminCourtHandler) CreateCourt(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Check if court information has created before....
		court, _ := ach.courtServ.Court()

		if court == nil {
			ach.tmpl.ExecuteTemplate(w, "admin.court.new.layout", nil)
			return
		}

		ach.tmpl.ExecuteTemplate(w, "login.layout", nil)
	}

	if r.Method == http.MethodPost {
		//Create a new court

		court_name := r.FormValue("court_name")
		court_level := r.FormValue("court_level")
		court_address := r.FormValue("court_address")
		court_phone := r.FormValue("court_phone")

		newCourt := entity.Court{CourtName: court_name, CourtLevel: court_level, CourtAddress: court_address, CourtPhone: court_phone}

		_, err2 := ach.courtServ.CreateCourt(&newCourt)

		if len(err2) > 0 {
			// panic(err2)
			ach.tmpl.ExecuteTemplate(w, "admin.court.new.layout", "Error Creating the Court")
		}

		ach.tmpl.ExecuteTemplate(w, "login.layout", nil)
	}
}

func (ach *AdminCourtHandler) UpdateCourt(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Check if court information has created before....
		//If there is get the court info for editing...
	}

	if r.Method == http.MethodPost {
		//update the court info

	}
}

func (ach *AdminCourtHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//Check if admin has created before....
		//Check if court information has created before....
		court, _ := ach.courtServ.Court()

		if court == nil {
			ach.tmpl.ExecuteTemplate(w, "admin.court.new.layout", nil)
			return
		}

		admin, _ := ach.courtServ.Admin()

		if admin == nil {
			ach.tmpl.ExecuteTemplate(w, "admin.account.new.layout", nil)
			return
		}

		ach.tmpl.ExecuteTemplate(w, "login.layout", nil)
	}

	if r.Method == http.MethodPost {
		//Create a new admin

		admin_id := r.FormValue("admin_id")
		admin_pwd := r.FormValue("admin_pwd")
		admin_pwd_conf := r.FormValue("admin_pwd_conf")

		if admin_pwd != admin_pwd_conf {
			ach.tmpl.ExecuteTemplate(w, "admin.account.new.layout", nil)
			return
		}

		hasher := md5.New()
		hasher.Write([]byte(admin_pwd))
		pwdnew := hex.EncodeToString(hasher.Sum(nil))

		newAdmin := entity.Admin{AdminId: admin_id, AdminPwd: pwdnew}

		_, err2 := ach.courtServ.CreateAdmin(&newAdmin)

		if len(err2) > 0 {
			// panic(err2)
			ach.tmpl.ExecuteTemplate(w, "admin.account.new.layout", "Error Creating Admin Account")
		}

		ach.tmpl.ExecuteTemplate(w, "login.layout", nil)

	}
}
