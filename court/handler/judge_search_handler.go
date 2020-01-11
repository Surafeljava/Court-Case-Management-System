package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/danieltefera/Project_Search/user"
)

//JudgeSearchHandler ...
type JudgeSearchHandler struct {
	judgeSearchService user.JudgeSearchService
}

//NewJudgeSearchHandler -
func NewJudgeSearchHandler(judgeSearch user.JudgeSearchService) *JudgeSearchHandler {
	return &JudgeSearchHandler{judgeSearchService: judgeSearch}
}

//Judges -
func (uh *JudgeSearchHandler) Judges(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, errs := uh.judgeSearchService.Judges()

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

//GetSingleJudge -
func (uh *JudgeSearchHandler) GetSingleJudge(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	//id, err := strconv.Atoi(ps.ByName("id"))

	idraw := r.FormValue("id")
	id, err := strconv.Atoi(idraw)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.judgeSearchService.Judge(uint(id))

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
