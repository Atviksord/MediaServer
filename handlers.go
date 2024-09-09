package main

import (
	"bytes"
	"fmt"
	"io"
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
func (cfg *apiconfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body %v", err)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		err = r.ParseForm()
		if err != nil {
			fmt.Printf("Error parsing form %v", err)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check DB for match, if match serve index. (make JWT etc for auth endpoints)
		_, err = cfg.db.Login(r.Context(), database.LoginParams{
			Username: username,
			Password: password})
		if err != nil {
			fmt.Println("No such user")
			http.ServeFile(w, r, "./static/login.html")
			return
		}

		http.ServeFile(w, r, "index.html")

	}

}
func (cfg *apiconfig) signupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body %v", err)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		err = r.ParseForm()
		if err != nil {
			fmt.Printf("Error parsing form %v", err)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check DB for match IF NOT EXIST create user in DB and automatically LOG IN (jwt creation, also hash password etc)
		_, err = cfg.db.GetUser(r.Context(), username)
		if err != nil {
			cfg.db.CreateUser(r.Context(), database.CreateUserParams{
				Username:  username,
				Password:  password,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC()})

			fmt.Printf("User does not exist, creating new user %s with the password %s", username, password)
			http.ServeFile(w, r, "./static/login.html")

			return
		}

	}

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
