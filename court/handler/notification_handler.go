package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		notif, _ := nh.notfService.ViewNotification(uint(id))

		nh.tmpl.ExecuteTemplate(w, "admin.notifications.update.layout", notif)

	} else if r.Method == http.MethodPost {

		id, _ := strconv.Atoi(r.FormValue("id"))
		notTitle := r.FormValue("not_title")
		notLevel := r.FormValue("not_level")
		notDescription := r.FormValue("not_desc")

		notification := entity.Notification{ID: uint(id), NotDescription: notDescription, NotTitle: notTitle, NotLevel: notLevel, NotDate: time.Now()}

		_, err := nh.notfService.UpdateNotification(&notification)

		if len(err) > 0 {
			fmt.Println(">> -- >> -- >> error on sending Notification to the updateNotification func")
			panic(err)
		}

		http.Redirect(w, r, "/admin/notifications", http.StatusSeeOther)

	} else {

		http.Redirect(w, r, "/admin/notifications", http.StatusSeeOther)

	}
}

// AdminNotifications handle requests on route /admin/notifications
func (nh *NotificationHandler) AdminNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, errs := nh.notfService.Notifications()
	if len(errs) > 0 {
		return
	}

	notificationStore := []entity.Notification{}

	for i := 0; i < len(notifications); i++ {

		notificationStore = append(notificationStore, notifications[i])

	}
	nh.tmpl.ExecuteTemplate(w, "admin.notifications.layout", notificationStore)

}

// AdminPostNotification handles POST coming from admin url request
func (nh *NotificationHandler) AdminPostNotification(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		//r.ParseForm()
		notTitle := r.FormValue("not_title")
		notLevel := r.FormValue("not_level")
		notDescription := r.FormValue("not_desc")

		notification := entity.Notification{NotDescription: notDescription, NotTitle: notTitle, NotLevel: notLevel, NotDate: time.Now()}
		_, errs := nh.notfService.PostNotification(&notification)

		if len(errs) > 0 {
			panic(errs)
		}

		http.Redirect(w, r, "/admin/notifications", http.StatusSeeOther)
	} else {
		nh.tmpl.ExecuteTemplate(w, "admin.notifications.postnotification.layout", nil)
	}

}

//AdminDeleteNotification existing Notification
func (nh *NotificationHandler) AdminDeleteNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		_, errs := nh.notfService.DeleteNotification(uint(id))

		if len(errs) > 0 {
			http.Redirect(w, r, "/admin/notifications", http.StatusSeeOther)
		}

		http.Redirect(w, r, "/admin/notifications", http.StatusSeeOther)

	}
}
