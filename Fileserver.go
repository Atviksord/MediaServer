package main

import (
	"html/template"
	"net/http"
)

func (cfg *apiconfig) templateInector(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Data to inject
	type data struct {
		Content string
	}
	datag := data{
		Content: "Inject testing",
	}

	// Execute the template with the data
	err = tmpl.Execute(w, datag)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

	http.ServeFile(w, r, "index.html")
}
