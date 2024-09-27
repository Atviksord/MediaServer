package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Atviksord/MediaServer/internal/database"
)

type MediaItem struct {
	Title    string
	FilePath string
	Format   string
}
type userInfo struct {
	Username string
}
type PageData struct {
	User   userInfo
	Title  string
	Videos []MediaItem
	Images []MediaItem
}

// Dynamic Injection of Data function.
func (cfg *apiconfig) templateInjector(w http.ResponseWriter, r *http.Request, user database.User) {
	// Parse the HTML templates
	tmpl, err := template.ParseFiles("index.html", "static/imageTemplate.html", "static/videoTemplate.html", "static/userDetailTemplate.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Data to inject will be generated from SQL database entries in the future.

	datag := PageData{
		User: userInfo{
			Username: user.Username,
		},
		Title: "Picture 1",
		Images: []MediaItem{
			{Title: "Picture 1", FilePath: "/Media/GOPHER.png"},
			{Title: "Picture 2", FilePath: "/Media/logo.png"},
			{Title: "Picture 3", FilePath: "/Media/primagen.jpg"}},
		Videos: []MediaItem{
			{Title: "Cool ducks running around", FilePath: "/static/Media/video1.mp4", Format: "video/mp4"},
		},
	}

	// Execute the template with the data
	err = tmpl.Execute(w, datag)
	if err != nil {
		http.Error(w, "Error rendering data template ddd", http.StatusInternalServerError)
		fmt.Println(err)
	}

}
