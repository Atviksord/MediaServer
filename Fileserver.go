package main

import (
	"context"
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
	// SQL get media from DB (NOT FULLY TESTED)
	_, err = cfg.sqlMediaGetter(user)
	if err != nil {
		fmt.Printf("Error getting Media data from SQL db %v", err)
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

func (cfg *apiconfig) sqlMediaGetter(user database.User) (PageData, error) {

	allMedia, err := cfg.db.GetAllMedia(context.Background())
	if err != nil {
		fmt.Println("Couldnt get media data from database")
	}
	trueData := PageData{}

	for _, datapoint := range allMedia {

		trueData.User.Username = user.Username

		if datapoint.Format == "video" {
			trueData.Videos = append(trueData.Videos, MediaItem{Title: datapoint.MediaName, FilePath: datapoint.FilePath, Format: datapoint.Format})

		}
		if datapoint.Format == "image" {
			trueData.Images = append(trueData.Images, MediaItem{Title: datapoint.MediaName, FilePath: datapoint.FilePath})

		}

	}

}
