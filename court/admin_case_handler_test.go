package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/caseUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	notificationUseRep "github.com/Surafeljava/Court-Case-Management-System/notificationUse/repository"
	notificationUseSERV "github.com/Surafeljava/Court-Case-Management-System/notificationUse/service"
)

func TestAllCases(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := notificationUseRep.NewMockNotificationRepo(nil)
	NotifServ := notificationUseSERV.NewNotificationServiceImpl(NotifRepo)

	CaHRepo := repository.NewMockCaseRepositoryImpl(nil)
	CaHServ := service.NewCaseServiceImpl(CaHRepo)

	adminCaseHandler := handler.NewCaseHandler(tmpl, CaHServ, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cases", adminCaseHandler.Cases)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/cases")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Killing two people")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestNewCase(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := notificationUseRep.NewMockNotificationRepo(nil)
	NotifServ := notificationUseSERV.NewNotificationServiceImpl(NotifRepo)

	CaHRepo := repository.NewMockCaseRepositoryImpl(nil)
	CaHServ := service.NewCaseServiceImpl(CaHRepo)

	adminCaseHandler := handler.NewCaseHandler(tmpl, CaHServ, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cases/new", adminCaseHandler.Cases)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("case_title", entity.CaseMock.CaseTitle)
	form.Add("case_desc", entity.CaseMock.CaseDesc)
	form.Add("case_type", entity.CaseMock.CaseType)
	form.Add("case_judge", entity.CaseMock.CaseJudge)

	resp, err := tc.PostForm(sURL+"/admin/cases/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Killing two people")) {
		t.Errorf("want body to contain %q", body)
	}
}

func TestUpdateCase(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := notificationUseRep.NewMockNotificationRepo(nil)
	NotifServ := notificationUseSERV.NewNotificationServiceImpl(NotifRepo)

	CaHRepo := repository.NewMockCaseRepositoryImpl(nil)
	CaHServ := service.NewCaseServiceImpl(CaHRepo)

	adminCaseHandler := handler.NewCaseHandler(tmpl, CaHServ, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cases/update", adminCaseHandler.Cases)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("id", string(entity.CaseMock.ID))
	form.Add("case_title", entity.CaseMock.CaseTitle)
	form.Add("case_desc", entity.CaseMock.CaseDesc)
	form.Add("case_type", entity.CaseMock.CaseType)
	form.Add("case_judge", entity.CaseMock.CaseJudge)

	resp, err := tc.PostForm(sURL+"/admin/cases/update?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Killing two people")) {
		t.Errorf("want body to contain %q", body)
	}
}

func TestDeleteCase(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := notificationUseRep.NewMockNotificationRepo(nil)
	NotifServ := notificationUseSERV.NewNotificationServiceImpl(NotifRepo)

	CaHRepo := repository.NewMockCaseRepositoryImpl(nil)
	CaHServ := service.NewCaseServiceImpl(CaHRepo)

	adminCaseHandler := handler.NewCaseHandler(tmpl, CaHServ, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/cases/delete", adminCaseHandler.Cases)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("id", string(entity.CaseMock.ID))
	form.Add("case_title", entity.CaseMock.CaseTitle)
	form.Add("case_desc", entity.CaseMock.CaseDesc)
	form.Add("case_type", entity.CaseMock.CaseType)
	form.Add("case_judge", entity.CaseMock.CaseJudge)

	resp, err := tc.PostForm(sURL+"/admin/cases/delete?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Killing two people")) {
		t.Errorf("want body to contain %q", body)
	}

}
