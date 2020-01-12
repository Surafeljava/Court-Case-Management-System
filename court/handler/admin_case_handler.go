package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse"
)

type CaseHandler struct {
	tmpl    *template.Template
	caseSrv caseUse.CaseService
}

func NewCaseHandler(T *template.Template, CS caseUse.CaseService) *CaseHandler {
	return &CaseHandler{tmpl: T, caseSrv: CS}
}

//Get all the cases in the court to the admin
func (lh *CaseHandler) Cases(w http.ResponseWriter, r *http.Request) {
	cases := []entity.Case{}
	cases, err := lh.caseSrv.Cases()
	if err != nil {
		panic(err)
	}
	lh.tmpl.ExecuteTemplate(w, "admin.cases.layout", cases)
}

//Add new Cases
func (lh *CaseHandler) NewCase(w http.ResponseWriter, r *http.Request) {
	//lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	//TODO Create a new case here...

	if r.Method == http.MethodPost {
		//TODO get the number of cases and generate a new case number
		//Get the case number
		allCases, err1 := lh.caseSrv.Cases()
		if err1 != nil {
			panic(err1)
		}

		cs_num := len(allCases) + 1
		case_num := fmt.Sprintf("CS%d", cs_num)

		case_title := r.FormValue("case_title")
		case_desc := r.FormValue("case_desc")
		case_type := r.FormValue("case_type")
		court_date := r.FormValue("court_date")
		case_judge := r.FormValue("case_judge")

		the_court_date, _ := time.Parse("2006-01-02", court_date)
		//the_case_creation, _ := time.Parse("2006-01-02", time.Now())

		newcs := entity.Case{CaseNum: case_num, CaseTitle: case_title, CaseDesc: case_desc, CaseStatus: "0", CaseType: case_type, CaseCreation: time.Now(), CaseCourtDate: the_court_date, CaseJudge: case_judge}

		err2 := lh.caseSrv.CreateCase(&newcs)

		if len(err2) > 0 {
			panic(err2)
		}

	} else {
		lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	}

	http.Redirect(w, r, "/admin/cases", http.StatusSeeOther)
}

//Delete existing cases
func (lh *CaseHandler) DeleteCase(w http.ResponseWriter, r *http.Request) {
	lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	//TODO Delete case here...
}

//Update existing cases
func (lh *CaseHandler) UpdateCase(w http.ResponseWriter, r *http.Request) {
	//lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	//TODO Update case here...
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		//idr := 30
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		cs, _ := lh.caseSrv.Case(id)

		lh.tmpl.ExecuteTemplate(w, "admin.case.update.layout", cs)

	} else if r.Method == http.MethodPost {

		id, _ := strconv.Atoi(r.FormValue("id"))
		case_num := r.FormValue("case_num")
		case_title := r.FormValue("case_title")
		case_desc := r.FormValue("case_desc")
		case_type := r.FormValue("case_type")
		case_creation := time.Now()
		case_court_date, _ := time.Parse("2006-01-02", r.FormValue("court_date"))
		case_judge := r.FormValue("case_judge")

		cs := entity.Case{ID: uint(id), CaseNum: case_num, CaseTitle: case_title, CaseDesc: case_desc, CaseType: case_type, CaseCreation: case_creation, CaseCourtDate: case_court_date, CaseJudge: case_judge}
		_, err := lh.caseSrv.UpdateCase(&cs)

		if len(err) > 0 {
			fmt.Println(">> -- >> -- >> error on sending the Case to the updateCase func")
			panic(err)
		}

		http.Redirect(w, r, "/admin/cases", http.StatusSeeOther)

	} else {

		http.Redirect(w, r, "/admin/cases", http.StatusSeeOther)
	}
}

//Update a case by extending the case court date >> by the judge
func (lh *CaseHandler) ExtendCase(w http.ResponseWriter, r *http.Request) {
	lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	//TODO Extend case court date here...
}

//Close a case by adding final decision and description >> by the judge
func (lh *CaseHandler) CloseCase(w http.ResponseWriter, r *http.Request) {
	lh.tmpl.ExecuteTemplate(w, "admin.newcase.layout", nil)
	///TODO Close case here...
}
