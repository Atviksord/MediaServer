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
		user, err := cfg.db.Login(r.Context(), database.LoginParams{
			Username: username,
			Password: password})
		if err != nil {
			fmt.Println("No such user or password")
			http.ServeFile(w, r, "./static/login.html")
			return
		}
		refreshToken, err := cfg.generateRandomToken()
		if err != nil {
			fmt.Println("Could not generate random token")
		}
		_, err = cfg.db.AddAccessToken(r.Context(), database.AddAccessTokenParams{
			Username:     username,
			Refreshtoken: sql.NullString{String: refreshToken, Valid: true}})
		if err != nil {
			fmt.Printf("Error generating random access token %v", err)
		}
		cfg.cookieFactory(w, refreshToken)
		cfg.templateInjector(w, r, user)

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

// Creates a cookie to store refreshtoken on login
func (cfg *apiconfig) cookieFactory(w http.ResponseWriter, refreshToken string) {
	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Println("Cookie made successfully")
}

func (cfg *apiconfig) logoutHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	if r.Method == "POST" {
		cfg.db.DelAccessToken(r.Context(), user.Refreshtoken)
		fmt.Printf("user has logged out %s", user.Username)

		http.ServeFile(w, r, "./static/login.html")

	}

}
