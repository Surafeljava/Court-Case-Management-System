package handler

import (
	"html/template"
	"net/http"

	user "github.com/Surafeljava/Court-Case-Management-System/SearchUse"
)

var tmpl = template.Must(template.ParseGlob("../UI/templates/*.html"))

//CaseSearchHandler ...
type CaseSearchHandler struct {
	caseSearchService user.CaseSearchService
}

//NewCaseSearchHandler -
func NewCaseSearchHandler(caseSearch user.CaseSearchService) *CaseSearchHandler {
	return &CaseSearchHandler{caseSearchService: caseSearch}
}

//Cases -
func (uh *CaseSearchHandler) Cases(w http.ResponseWriter, r *http.Request) {
	cases, errs := uh.caseSearchService.Cases()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmpl.ExecuteTemplate(w, "caseSearchResult.layout", cases)
	return
}

//GetSingleCase retrives a SINGLE CASE given the Case Number
func (uh *CaseSearchHandler) GetSingleCase(w http.ResponseWriter,
	r *http.Request) {

	caseNum := r.FormValue("search_caseNum")

	singlecase, errs := uh.caseSearchService.Case(caseNum)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	tmpl.ExecuteTemplate(w, "caseSearchResultSingle.layout", singlecase)
	return
}

//TODO Input Validation
