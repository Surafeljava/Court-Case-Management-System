package handler

import (
	"html/template"
	"net/http"

	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
)

var tmplJudge = template.Must(template.ParseGlob("../UI/templates/*.html"))

//JudgeSearchHandler ...
type JudgeSearchHandler struct {
	judgeSearchService user.JudgeSearchService
}

//NewJudgeSearchHandler -
func NewJudgeSearchHandler(judgeSearch user.JudgeSearchService) *JudgeSearchHandler {
	return &JudgeSearchHandler{judgeSearchService: judgeSearch}
}

//Judges -
func (uh *JudgeSearchHandler) Judges(w http.ResponseWriter, r *http.Request) {
	judges, errs := uh.judgeSearchService.Judges()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	tmplJudge.ExecuteTemplate(w, "judgeSearchResult.layout", judges)
	return
}

//GetSingleJudge -
func (uh *JudgeSearchHandler) GetSingleJudge(w http.ResponseWriter,
	r *http.Request) {

	judgeID := r.FormValue("search_judgeID")

	judge, errs := uh.judgeSearchService.Judge(judgeID)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	tmplJudge.ExecuteTemplate(w, "judgeSearchResultSingle.layout", judge)
	return
}
