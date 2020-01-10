package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Surafeljava/Court-Case-Management-System/entity"
	"github.com/bb/Court-Case-Management-System/notificationUse"
	"github.com/julienschmidt/httprouter"
)

//NotificationHandler
type NotificationHandler struct {
	notfService notificationUse.NotificationService
}

// PostNotification handles POST coming from admin url request
func (nh *NotificationHandler) PostNotifications(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	notf := &entity.Notification{}

	err := json.Unmarshal(body, notf)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	notifications, errs := nh.notfService.PostNotification(notf)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//p := fmt.Sprintf("/v1/admin/notifications/%s", notf.NotfTitle)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}
