package handlers

import (
	"app/sessions"
	"app/templates"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	sess, _ := r.Context().Value("sess").(*sessions.Session)
	t := templates.GetLayoutTemplate("home", "/templates/htmls/layout.html", "/templates/htmls/home.html")
	data := templates.TemplateData{
		Title: "Home",
		Data:  sess.Data,
	}
	t.ExecuteTemplate(w, "layout", data)
}
