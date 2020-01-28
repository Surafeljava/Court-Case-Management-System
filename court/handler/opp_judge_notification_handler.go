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

//CheckForWhom ..it checks for whom the notification will be posted
func CheckForWhom(level string) int {

	fmt.Println(level)
	if level == "all" {
		// fmt.Println(">>>> AD-checked! and 0 returned")
		return 0
	} else if level == "judges" {
		return 1
	} else if level == "opponents" {
		return 2
	}
	return -1
}

//DeleteNotification deletes notification by id
func (ojh *OppJNotificationHandler) DeleteNotification(w http.ResponseWriter, r *http.Request) {

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

			_, errs := ojh.notfService.DeleteNotification(uint(id))

			if len(errs) > 0 {
				return
			}

			http.Redirect(w, r, "/opponent/notifications", http.StatusSeeOther)
		} else if forWhom == 0 {
			//display error message
			//if notification level is "all" cant be deleted
		}

	}
}

//NotificationsJudge retrieves all notification of judges
func (ojh *OppJNotificationHandler) NotificationsJudge(w http.ResponseWriter, r *http.Request) {
	notifications, errs := ojh.notfService.Notifications()
	if len(errs) > 0 {
		return
	}

	notificationStore := []entity.Notification{}

	for i := 0; i < len(notifications); i++ {

		forWhom := CheckForWhom(notifications[i].NotLevel)

		if forWhom == 0 || forWhom == 1 {
			notificationStore = append(notificationStore, notifications[i])

		}

	}
	ojh.tmpl.ExecuteTemplate(w, "judge.notifications.layout", notificationStore)

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
