package handlers

import (
	"app/sessions"
	"app/templates"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("sess").(*sessions.Session)
	t := templates.GetTemplate("/templates/htmls/home.html")
	t.Execute(w, sess.Data)
}
