package handlers

import (
	"app/models"
	"app/sessions"
	"app/templates"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	view_data := struct {
		ErrorMessages []string
		Name          string
		Pass          string
	}{
		[]string{},
		"",
		"",
	}
	if r.Method == "POST" {
		view_data.Name = r.FormValue("name")
		view_data.Pass = r.FormValue("pass")
		user := models.GetUserByNamePass(view_data.Name, view_data.Pass)
		if user != nil {
			sessions.New(r, w, map[string]interface{}{"usr_id": user.ID, "usr_name": user.Name})
			// redirect
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			view_data.ErrorMessages = append(view_data.ErrorMessages, "Invalid user or password")
		}
	}
	data := templates.TemplateData{
		Title: "login",
		Data:  view_data,
	}
	t := templates.GetLayoutTemplate("login", "/templates/htmls/layout.html", "/templates/htmls/login.html")
	t.ExecuteTemplate(w, "layout", data)
}
