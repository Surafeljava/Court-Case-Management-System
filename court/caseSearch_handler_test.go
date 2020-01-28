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

func TestCases(t *testing.T) {

	caseRepo := repository.NewMockCaseSearchRepo(nil)
	caseServ := service.NewCaseSearchService(caseRepo)

	caseSearchHandler := handler.NewCaseSearchHandler(caseServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/admin/cases", caseSearchHandler.Cases)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/cases")
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

//------------------------------------------------------
// Below is FOR TESTING SINGLE CASE SEARCH FUNCTION!!! |
//------------------------------------------------------

func TestGetSingleCase(t *testing.T) {

	caseRepo := repository.NewMockCaseSearchRepo(nil)
	caseServ := service.NewCaseSearchService(caseRepo)

	caseSearchHandler := handler.NewCaseSearchHandler(caseServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/admin/cases/singlecase", caseSearchHandler.GetSingleCase)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	resp, err := tc.Get(sURL + "/v1/admin/cases/singlecase?search_caseNum=CS1")
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
