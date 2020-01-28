package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	entity "github.com/Surafeljava/Court-Case-Management-System/Entity"

	notificationUse "github.com/Surafeljava/Court-Case-Management-System/notificationUse"
)

//JudgeNotificationHandler struct opponent and judge notification handler
type JudgeNotificationHandler struct {
	tmpl        *template.Template
	notfService notificationUse.NotificationService
}

//NewJudgeNotificationHandler creates Notification handler
func NewJudgeNotificationHandler(T *template.Template, NotfServ notificationUse.NotificationService) *JudgeNotificationHandler {
	return &JudgeNotificationHandler{tmpl: T, notfService: NotfServ}

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

//NotificationsJudge retrieves all notification of judges
func (juh *JudgeNotificationHandler) NotificationsJudge(w http.ResponseWriter, r *http.Request) {
	notifications, errs := juh.notfService.Notifications()
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
	juh.tmpl.ExecuteTemplate(w, "judge.notifications.layout", notificationStore)

}

//SingleNotificationJudge retrieves single notification of judges
func (juh *JudgeNotificationHandler) SingleNotificationJudge(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		//idr := 30
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		notiff, _ := juh.notfService.ViewNotification(uint(id))

		juh.tmpl.ExecuteTemplate(w, "opponent.notifications.update.layout", notiff)

	} else if r.Method == http.MethodPost {
		http.Redirect(w, r, "/judge/notifications", http.StatusSeeOther)
	}

}

//DeleteJudgeNotification deletes notification by id
func (juh *JudgeNotificationHandler) DeleteJudgeNotification(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			fmt.Println("Can't get the id!")
			panic(err)
		}

		notification, errs := juh.notfService.ViewNotification(uint(id))
		if len(errs) > 0 {
			return
		}

		forWhom := CheckForWhom(notification.NotLevel)
		if forWhom == 1 {

			_, errs := juh.notfService.DeleteNotification(uint(id))

			if len(errs) > 0 {
				http.Redirect(w, r, "/judge/notifications", http.StatusSeeOther)
			}

			http.Redirect(w, r, "/judge/notifications", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/judge/notifications", http.StatusSeeOther)
		}
	}
}
