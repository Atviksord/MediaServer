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
	"golang.org/x/crypto/bcrypt"
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

		firstUser, err := cfg.db.GetUser(r.Context(), username)
		if err != nil {
			fmt.Println("No such username", err)
			http.ServeFile(w, r, "./static/login.html")
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(firstUser.Password), []byte(password))
		if err == nil {
			user, err := cfg.db.Login(r.Context(), database.LoginParams{
				Username: username,
				Password: firstUser.Password})
			if err != nil {
				fmt.Println("No such user or password")
				http.ServeFile(w, r, "./static/login.html")
				return
			}
			refreshToken, err := cfg.generateRandomToken()
			if err != nil {
				fmt.Println("Could not generate random token")
				http.ServeFile(w, r, "./static/login.html")
				return

			}
			_, err = cfg.db.AddAccessToken(r.Context(), database.AddAccessTokenParams{
				Username:     username,
				Refreshtoken: sql.NullString{String: refreshToken, Valid: true}})
			if err != nil {
				fmt.Printf("Error generating random access token %v", err)
				http.ServeFile(w, r, "./static/login.html")
				return
			}
			cfg.cookieFactory(w, refreshToken)
			cfg.templateInjector(w, r, user)
			return

		} else {
			fmt.Println("No such user or wrong password")
			http.ServeFile(w, r, "./static/login.html")
			return
		}

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
