package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/Atviksord/MediaServer/internal/database"
)

type UserInfo struct {
	Username   string
	password   string
	updated_at time.Time
	created_At time.Time
}

func (cfg *apiconfig) startingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	// AUTH CHECK (check if already logged in)

	// Serve Server login if not API auth'd or JWT
	http.ServeFile(w, r, "./static/login.html")

}

func (cfg *apiconfig) authWrapper(w http.ResponseWriter, r *http.Request, user database.User) {

}

func (cfg *apiconfig) handlerRegistry(mux *http.ServeMux) {
	// Fileserver Handler creation
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Send to loginpage/root check
	mux.HandleFunc("/", cfg.startingHandler)
	mux.HandleFunc("POST /login", cfg.loginHandler)
	mux.HandleFunc("POST /signup", cfg.signupHandler)

}
