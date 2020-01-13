package handler

import (
	"fmt"
	"html/template"
	"net/http"

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

	Example := cases.CaseNum + " " + "Opponent ID : " + opp.OppName + "Witness:" + wit.WitnessDoc + "Decision: " + dic.DecisionDesc
	fmt.Println(oppNum)
	fmt.Println(Example)

	if len(err) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmplAppeal.ExecuteTemplate(w, "appeal.html", Example)
	return
}

//OppTrial is only for trial
func (oppAp *AppealHandler) OppTrial(w http.ResponseWriter, r *http.Request) {
	tmplAppeal.ExecuteTemplate(w, "opponent.home.layout", nil)
	return
}
