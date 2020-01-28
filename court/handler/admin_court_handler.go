package handler

import (
	"html/template"
	"net/http"

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

	}

	if r.Method == http.MethodPost {
		//Create a new admin

	}
}
