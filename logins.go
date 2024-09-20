package main

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

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
			fmt.Println("No such user or password")
			http.ServeFile(w, r, "./static/login.html")
			return
		}
		randomAccess, err := cfg.generateRandomToken()
		if err != nil {
			fmt.Println("Could not generate random token")
		}
		_, err = cfg.db.AddAccessToken(r.Context(), database.AddAccessTokenParams{
			Username:     username,
			Refreshtoken: sql.NullString{String: randomAccess, Valid: true}})
		if err != nil {
			fmt.Printf("Error generating random access token %v", err)
		}
		cfg.templateInjector(w, r)

	}

}

func (cfg *apiconfig) generateRandomToken() (string, error) {
	b := make([]byte, 15)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (cfg *apiconfig) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("TEST TEST")
	}

}
