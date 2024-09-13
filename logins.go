package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/Atviksord/MediaServer/internal/database"
)

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

	}

}
