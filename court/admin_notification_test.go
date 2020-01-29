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
	"github.com/Surafeljava/Court-Case-Management-System/court/handler"
	"github.com/Surafeljava/Court-Case-Management-System/notificationUse/repository"
	"github.com/Surafeljava/Court-Case-Management-System/notificationUse/service"
)

func TestAdminNotifications(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := repository.NewMockNotificationRepo(nil)
	NotifServ := service.NewNotificationServiceImpl(NotifRepo)

	adminNotifHandler := handler.NewNotificationHandler(tmpl, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/notifications", adminNotifHandler.AdminNotifications)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/notifications")
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

	if !bytes.Contains(body, []byte("Team meeting")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminPostNotification(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := repository.NewMockNotificationRepo(nil)
	NotifServ := service.NewNotificationServiceImpl(NotifRepo)

	adminNotifHandler := handler.NewNotificationHandler(tmpl, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/notifications/postnotification", adminNotifHandler.AdminNotifications)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("not_title", entity.NotificationMock.NotTitle)
	form.Add("not_level", entity.NotificationMock.NotLevel)
	form.Add("not_desc", entity.NotificationMock.NotDescription)

	resp, err := tc.PostForm(sURL+"/admin/notifications/postnotification", form)
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

	if !bytes.Contains(body, []byte("Team meeting")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminUpdateNotification(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := repository.NewMockNotificationRepo(nil)
	NotifServ := service.NewNotificationServiceImpl(NotifRepo)

	adminNotifHandler := handler.NewNotificationHandler(tmpl, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/notifications/update", adminNotifHandler.AdminNotifications)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("id", string(entity.NotificationMock.ID))
	form.Add("not_title", entity.NotificationMock.NotTitle)
	form.Add("not_level", entity.NotificationMock.NotLevel)
	form.Add("not_desc", entity.NotificationMock.NotDescription)

	resp, err := tc.PostForm(sURL+"/admin/notifications/update?id=1", form)
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

	if !bytes.Contains(body, []byte("Team meeting")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminCategoresDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../UI/templates/*"))

	NotifRepo := repository.NewMockNotificationRepo(nil)
	NotifServ := service.NewNotificationServiceImpl(NotifRepo)

	adminNotifHandler := handler.NewNotificationHandler(tmpl, NotifServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/notifications/delete", adminNotifHandler.AdminNotifications)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("id", string(entity.NotificationMock.ID))
	form.Add("not_title", entity.NotificationMock.NotTitle)
	form.Add("not_level", entity.NotificationMock.NotLevel)
	form.Add("not_desc", entity.NotificationMock.NotDescription)

	resp, err := tc.PostForm(sURL+"/admin/notifications/delete?id=1", form)
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

	if !bytes.Contains(body, []byte("Team meeting")) {
		t.Errorf("want body to contain %q", body)
	}

}
