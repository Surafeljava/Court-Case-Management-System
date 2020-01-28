package handler

import (
	"html/template"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/reportUse"
)

type ReportHandler struct {
	tmpl    *template.Template
	repServ reportUse.ReportService
}

func NewReportHandler(T *template.Template, rs reportUse.ReportService) *ReportHandler {
	return &ReportHandler{tmpl: T, repServ: rs}
}

func (rh *ReportHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	closedCases, er1 := rh.repServ.GetClosedCases()
	openCases, er2 := rh.repServ.GetOpenCases()
	allCases, er3 := rh.repServ.GetAllCases()

	allJudges, er4 := rh.repServ.GetAllJudges()
	criminalJudges, er5 := rh.repServ.GetCriminalJudges()
	civilJudges, er6 := rh.repServ.GetCivilJudges()

	allNotifications, er7 := rh.repServ.GetAllNotifications()

	if er1 != nil || er2 != nil || er3 != nil || er4 != nil || er5 != nil || er6 != nil || er7 != nil {
		return
	}

	data := struct {
		ClosedCases      int
		OpenCases        int
		AllCases         int
		AllJudges        int
		CriminalJudges   int
		CivilJudges      int
		AllNotifications int
	}{
		ClosedCases:      len(closedCases),
		OpenCases:        len(openCases),
		AllCases:         len(allCases),
		AllJudges:        len(allJudges),
		CriminalJudges:   len(criminalJudges),
		CivilJudges:      len(civilJudges),
		AllNotifications: len(allNotifications),
	}

	rh.tmpl.ExecuteTemplate(w, "admin.report.layout", data)

}
