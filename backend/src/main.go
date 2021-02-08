package main

import (
	"app/models"
	app_session "app/session"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	Port int = 80
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", accessLogMiddleware(loginCheckMiddleware(home)))
	http.HandleFunc("/login", accessLogMiddleware(login))
	http.HandleFunc("/logout", accessLogMiddleware(logout))
	log.Printf("Start Server on port:%d", Port)
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home")

	sess := r.Context().Value("sess")
	b, _ := json.MarshalIndent(sess, "", "\t")
	fmt.Fprint(w, fmt.Sprintf("\nlogged in: \n%s", string(b)))
}

func login(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	pass := r.FormValue("pass")
	user := models.GetUserByIDPass(name, pass)
	if user == nil {
		fmt.Fprint(w, "login failed")
	} else {
		app_session.New(r, w, map[string]interface{}{"usr_id": user.ID, "user_name": user.Name})
		fmt.Fprint(w, "logged in")
		fmt.Fprint(w, "\nuser: %#v", user)
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	app_session.Delete(r, w)
	fmt.Fprint(w, "logged out")
}

func accessLogMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func loginCheckMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if is_logged_in, sess := app_session.CheckSession(r); is_logged_in {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "sess", sess)
			r = r.WithContext(ctx)
			next(w, r)
		} else {
			fmt.Fprint(w, "required login")
		}
	}
}
