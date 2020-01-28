package handler

import (
	"html/template"
	"net/http"
	"time"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	"github.com/Surafeljava/Court-Case-Management-System/notificationUse"
)

//NotificationHandler struct
type NotificationHandler struct {
	tmpl        *template.Template
	notfService notificationUse.NotificationService
}

//NewNotificationHandler creates Notification handler
func NewNotificationHandler(T *template.Template, NotfServ notificationUse.NotificationService) *NotificationHandler {
	return &NotificationHandler{tmpl: T, notfService: NotfServ}
}

//AdminUpdateNotification updates notification
func (nh *NotificationHandler) AdminUpdateNotification(w http.ResponseWriter, r *http.Request) {
	//To Do ..|
	//update notification goes here
}

// AdminPostNotification handles POST coming from admin url request
func (nh *NotificationHandler) AdminPostNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		notTitle := r.FormValue("not_title")
		notLevel := r.FormValue("not_level")
		notDescription := r.FormValue("not_desc")

		notification := entity.Notification{NotDescription: notDescription, NotTitle: notTitle, NotLevel: notLevel, NotDate: time.Now()}
		_, errs := nh.notfService.PostNotification(&notification)

		if len(errs) > 0 {
			return
		}
		http.Redirect(w, r, "/admin/home", http.StatusSeeOther)

	} else {
		nh.tmpl.ExecuteTemplate(w, "admin.postNotification.html", nil)
	}

}
