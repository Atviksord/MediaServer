package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"

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
	datag, err := cfg.sqlMediaGetter(user)
	fmt.Println(datag)
	if err != nil {
		fmt.Printf("Error getting Media data from SQL db %v", err)
	}

	// Execute the template with the data
	err = tmpl.Execute(w, datag)
	if err != nil {
		http.Error(w, "Error rendering data template ddd", http.StatusInternalServerError)
		fmt.Println(err)
	}

}

func (cfg *apiconfig) sqlMediaGetter(user database.User) (PageData, error) {
	// Will make a custom page for favourited/followed videos.
	allMedia, err := cfg.db.GetAllMedia(context.Background())
	if err != nil {
		fmt.Println("Couldnt get media data from database")
	}
	trueData := PageData{}
	fmt.Println(allMedia[0].Format)

	for _, datapoint := range allMedia {

		trueData.User.Username = user.Username

		if datapoint.MediaType == "video" {
			trueData.Videos = append(trueData.Videos, MediaItem{Title: datapoint.MediaName, FilePath: datapoint.FilePath, Format: datapoint.Format})

		}
		if datapoint.MediaType == "image" {
			imagePath := strings.TrimPrefix(datapoint.FilePath, "static")
			trueData.Images = append(trueData.Images, MediaItem{Title: datapoint.MediaName, FilePath: imagePath})

		}

	}
	return trueData, nil

}
