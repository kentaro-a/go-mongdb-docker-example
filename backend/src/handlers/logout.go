package handlers

import (
	"app/sessions"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	sessions.Delete(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
