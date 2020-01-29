package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Surafeljava/Court-Case-Management-System/SearchUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/SearchUse/service"
	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
)

func TestJudges(t *testing.T) {

	judgeRepo := repository.NewMockJudgeSearchRepo(nil)
	judegServ := service.NewJudgeSearchService(judgeRepo)

	judgeSearchHandler := handler.NewJudgeSearchHandler(judegServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/admin/judges", judgeSearchHandler.Judges)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/judges")
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

	if !bytes.Contains(body, []byte("JU1")) {
		t.Errorf("want body to contain %q", body)
	}

}

//------------------------------------------------------
//FOR TESTING SINGLE JUDGE SEARCH FUNCTION!!!           |
//------------------------------------------------------

func TestGetSingleJudge(t *testing.T) {

	judgeRepo := repository.NewMockJudgeSearchRepo(nil)
	judegServ := service.NewJudgeSearchService(judgeRepo)

	judgeSearchHandler := handler.NewJudgeSearchHandler(judegServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/admin/judges/singlejudge", judgeSearchHandler.GetSingleJudge)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/judges/singlejudge?search_judgeID=JU1")
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

	if !bytes.Contains(body, []byte("JU1")) {
		t.Errorf("want body to contain %q", body)
	}

}
