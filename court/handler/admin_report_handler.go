package handler

import (
	"html/template"

	"github.com/Surafeljava/Court-Case-Management-System/reportUse"
)

type ReportHandler struct {
	tmpl    *template.Template
	repServ reportUse.ReportService
}

func NewReportHandler(T *template.Template, rs reportUse.ReportService) *ReportHandler {
	return &ReportHandler{tmpl: T, repServ: rs}
}
