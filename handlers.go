package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiconfig) loginServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	// Servel login if not API auth'd or JWT
	http.ServeFile(w, r, "./static/login.html")
	// test
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error biatch")
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println(username, password)
	}

}

func (cfg *apiconfig) authenticatedFileServer(w http.ResponseWriter, r *http.Request) {

}

func (cfg *apiconfig) handlerRegistry(mux *http.ServeMux) {
	// Fileserver Handler creation
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Send to loginpage/root check
	mux.HandleFunc("/", cfg.loginServer)

	//mux.HandleFunc("GET /", cfg.fileServer)

}
