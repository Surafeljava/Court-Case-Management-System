package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/Surafeljava/Court-Case-Management-System/notificationUse"
	"github.com/julienschmidt/httprouter"
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

//JudgeViewNotification handles GET coming from judge url request
func (nh *NotificationHandler) JudgeViewNotification(w http.ResponseWriter, r *http.Request) {

}

// AdminPostNotification handles POST coming from admin url request
func (nh *NotificationHandler) AdminPostNotification(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		l := r.ContentLength
		body := make([]byte, l)
		r.Body.Read(body)
		notff := entity.Notification{}
		notff.NotTitle = r.FormValue("NotTitle")
		notff.NotLevel = r.FormValue("NotLevel")
		notff.NotDescription = r.FormValue("NotDesc")
		notff.NotDate = time.Now()
		err := json.Unmarshal(body, notff)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		_, errs := nh.notfService.PostNotification(&notff)

		if len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {
		nh.tmpl.ExecuteTemplate(w, "admin.postNotification.html", nil)
	}

}
