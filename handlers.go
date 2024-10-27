package main

import (
	"bytes"
	"context"
	"database/sql"
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

// CUSTOM TYPE FOR HANDLERS THAT REQUIRE AUTH
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiconfig) startingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	// AUTH CHECK (check if already logged in)
	// Get cookie
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		fmt.Printf("Error getting cookie %v", err)

	} else {
		refreshToken := cookie.Value
		user, err := cfg.db.GetUserByToken(r.Context(), sql.NullString{String: refreshToken, Valid: true})
		if err != nil {
			fmt.Println("Couldnt get user by token from main handler")

		} else {
			cfg.templateInjector(w, r, user)
			return
		}

	}

	// Serve Server login if not API auth'd or JWT
	http.ServeFile(w, r, "./static/login.html")

}
func (cfg *apiconfig) searchHandler(w http.ResponseWriter, r *http.Request, user database.User) {
	// Get from post request the search term, then run a SQL search query using that keyword and serve only files that have that search term in the title
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
		searchTerm := r.FormValue("searchContent")
		// sql query to get searched files
		searchedFiles, err := cfg.db.GetSearchedMedia(context.Background(), sql.NullString{String: searchTerm, Valid: true})
		if err != nil {
			fmt.Println("Couldnt get searched results from DB")
		}

		// gets PageData to inject
		trueData, err := cfg.pageDataArranger(searchedFiles, user)
		if err != nil {
			fmt.Println("Couldnt get searchedfile data from pageDataArranger")

		}
		// Get value from search and inject the correct files

		cfg.searchedTemplateInjector(w, r, user, trueData)

	}

}

func (cfg *apiconfig) authWrapper(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get cookie
		cookie, err := r.Cookie("refreshToken")
		if err != nil {
			fmt.Printf("Error getting cookie %v", err)
			return
		}
		refreshToken := cookie.Value

		// Get user from DB (check cookie for match)
		user, err := cfg.db.GetUserByToken(r.Context(), sql.NullString{String: refreshToken, Valid: true})
		if err != nil {
			fmt.Println("Didnt manage to get user by refreshToken")
			return
		}

		// Pass control from AUTH middleware to main handler
		handler(w, r, user)

	}
}

func (cfg *apiconfig) handlerRegistry(mux *http.ServeMux) {
	// Fileserver Handler creation
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Send to loginpage/root check
	mux.HandleFunc("/", cfg.startingHandler)
	mux.HandleFunc("POST /login", cfg.loginHandler)
	mux.HandleFunc("POST /signup", cfg.signupHandler)
	mux.HandleFunc("POST /logout", cfg.authWrapper(cfg.logoutHandler))
	mux.HandleFunc("POST /search", cfg.authWrapper(cfg.searchHandler))
	mux.HandleFunc("POST /togglefavourite", cfg.authWrapper(cfg.togglefavourite))
	mux.HandleFunc("GET /favouritelist", cfg.authWrapper(cfg.favouriteListHandler))

}
