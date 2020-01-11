package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/danieltefera/Project_Search/user"
)

//CaseSearchHandler ...
type CaseSearchHandler struct {
	caseSearchService user.CaseSearchService
}

//NewCaseSearchHandler -
func NewCaseSearchHandler(caseSearch user.CaseSearchService) *CaseSearchHandler {
	return &CaseSearchHandler{caseSearchService: caseSearch}
}

//Cases -
func (uh *CaseSearchHandler) Cases(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, errs := uh.caseSearchService.Cases()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

//GetSingleCase -
func (uh *CaseSearchHandler) GetSingleCase(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	//id, err := strconv.Atoi(ps.ByName("id"))

	idraw := r.FormValue("idc")
	id, err := strconv.Atoi(idraw)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.caseSearchService.Case(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
