package handler

import (
	"html/template"
	"net/http"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	user "github.com/Surafeljava/Court-Case-Management-System/appealUse"
)

var tmplAppeal = template.Must(template.ParseGlob("../UI/templates/*.html"))

//AppealHandler ...
type AppealHandler struct {
	appealService user.AppealService
}

//NewAppealHandler -
func NewAppealHandler(appeal user.AppealService) *AppealHandler {
	return &AppealHandler{appealService: appeal}
}

//OppAppeal -
func (oppAp *AppealHandler) OppAppeal(w http.ResponseWriter, r *http.Request) {
	oppNum := r.FormValue("opp_appeal")
	cases, opp, wit, dic, err := oppAp.appealService.Appeal(oppNum)

	data := entity.AppealForm{
		CaseNum:          cases.CaseNum,
		CaseCreationDate: cases.CaseCreation,
		CaseTitle:        cases.CaseTitle,
		CaseDesc:         cases.CaseDesc,
		OppName:          opp.OppName,
		OppGender:        opp.OppGender,
		OppAddress:       opp.OppAddress,
		OppPhone:         opp.OppPhone,
		WitDocm:          wit.WitnessDoc,
		WitTy:            wit.WitnessType,
		Decision:         dic.Decision,
		DecDate:          dic.DecisionDate,
		DacDesc:          dic.DecisionDesc,
	}

	if len(err) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	tmplAppeal.ExecuteTemplate(w, "appeal.layout", data)
	return
}

//OppTrial is only for trial
func (oppAp *AppealHandler) OppTrial(w http.ResponseWriter, r *http.Request) {
	tmplAppeal.ExecuteTemplate(w, "opponent.home.layout", nil)
	return
}
