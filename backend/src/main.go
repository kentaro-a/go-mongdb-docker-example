package main

import (
	"app/handlers"
	"app/middlewares"
	"fmt"
	"log"
	"net/http"
)

const (
	Port int = 80
)

//go:generate go-assets-builder -p templates -o templates/assets.go templates/htmls

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.HandleFunc("/", middlewares.AccessLogMiddleware(
		middlewares.LoginCheckMiddleware(handlers.Home),
	))
	http.HandleFunc("/login", middlewares.AccessLogMiddleware(handlers.Login))
	http.HandleFunc("/logout", middlewares.AccessLogMiddleware(handlers.Logout))
	log.Printf("Start Server on port:%d", Port)
	http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
}
