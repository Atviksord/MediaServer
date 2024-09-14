package main

import (
	"html/template"
	"net/http"
)

type MediaItem struct {
	Title    string
	FilePath string
	Format   string
}

// Dynamic Injection of Data function.
func (cfg *apiconfig) templateInjector(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML templates
	tmpl, err := template.ParseFiles("index.html", "static/imageTemplate.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Data to inject will be generated from SQL database entries in the future.
	type data struct {
		Title  string
		Videos []MediaItem
		Images []MediaItem
	}
	datag := data{
		Title: "Picture 1",
		Images: []MediaItem{
			{Title: "Picture 1", FilePath: "/Media/GOPHER.png"},
			{Title: "Picture 2", FilePath: "/Media/logo.png"},
			{Title: "Picture 3", FilePath: "/Media/primagen.jpg"}},
	}

	// Execute the template with the data
	err = tmpl.Execute(w, datag)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}
