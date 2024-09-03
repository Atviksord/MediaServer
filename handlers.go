package main

import (
	"net/http"
	"os"
)

func (cf *apiconfig) fileServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Could not load index.html", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write(data)

}

func (cfg *apiconfig) handlerRegistry(mux *http.ServeMux) {

	mux.HandleFunc("GET /", cfg.fileServer)

}
