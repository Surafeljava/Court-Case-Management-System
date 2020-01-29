package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"
	notificationUse "github.com/Surafeljava/Court-Case-Management-System/notificationUse"
)

//OppJNotificationHandler struct opponent and judge notification handler
type OppJNotificationHandler struct {
	tmpl        *template.Template
	notfService notificationUse.NotificationService
}

//NewOppJNotificationHandler creates Notification handler
func NewOppJNotificationHandler(T *template.Template, NotfServ notificationUse.NotificationService) *OppJNotificationHandler {
	return &OppJNotificationHandler{tmpl: T, notfService: NotfServ}

}

//DeleteOpponentNotification deletes notification by id
func (ojh *OppJNotificationHandler) DeleteOpponentNotification(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		notification, errs := ojh.notfService.ViewNotification(uint(id))
		if len(errs) > 0 {
			return
		}

		forWhom := CheckForWhom(notification.NotLevel)
		if forWhom == 1 {

			_, errs := ojh.notfService.DeleteNotification(uint(id))

			if len(errs) > 0 {
				return
			}

			http.Redirect(w, r, "/judge/notifications", http.StatusSeeOther)
		} else if forWhom == 2 {
			if forWhom == 2 {

				_, errs := ojh.notfService.DeleteNotification(uint(id))

				if len(errs) > 0 {
					return
				}

				http.Redirect(w, r, "/opponent/notifications", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/opponent/notifications", http.StatusSeeOther)
				//display error message
				//if notification level is "all" cant be deleted
			}

		}
	}
}

//NotificationsOpponent retrieves all notification of opponents
func (ojh *OppJNotificationHandler) NotificationsOpponent(w http.ResponseWriter, r *http.Request) {

	notifications, errs := ojh.notfService.Notifications()
	if len(errs) > 0 {
		return
	}

	notificationStore := []entity.Notification{}

	for i := 0; i < len(notifications); i++ {

		forWhom := CheckForWhom(notifications[i].NotLevel)

		if forWhom == 0 || forWhom == 2 {
			notificationStore = append(notificationStore, notifications[i])

		}

	}
	ojh.tmpl.ExecuteTemplate(w, "opponent.notifications.layout", notificationStore)

}

//SingleNotificationOpponent retrives single notification
func (ojh *OppJNotificationHandler) SingleNotificationOpponent(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		//idr := 30
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		notiff, _ := ojh.notfService.ViewNotification(uint(id))

		ojh.tmpl.ExecuteTemplate(w, "opponent.notifications.update.layout", notiff)

	} else if r.Method == http.MethodPost {
		http.Redirect(w, r, "/opponent/notifications", http.StatusSeeOther)
	} else {

		http.Redirect(w, r, "/opponent/notifications", http.StatusSeeOther)
	}
}
