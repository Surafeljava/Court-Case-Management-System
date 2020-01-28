package handler

import (
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type AdminCourtHandler struct {
	tmpl  *template.Template
	court caseUse.CourtService
}

func NewAdminCourtHandler(T *template.Template, JS caseUse.CourtService) *AdminCourtHandler {
	return &AdminCourtHandler{tmpl: T, court: JS}
}

func (ach *AdminCourtHandler) CreateCourt(w http.ResponseWriter, r *http.Request) {

}

func (ach *AdminCourtHandler) UpdateCourt(w http.ResponseWriter, r *http.Request) {

}

func (ach *AdminCourtHandler) AdminCreate(w http.ResponseWriter, r *http.Request) {

}
