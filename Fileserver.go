package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
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
	// SQL get media from DB
	datag, err := cfg.sqlMediaGetter(user)

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

// Gets all media with SQL query
func (cfg *apiconfig) sqlMediaGetter(user database.User) (PageData, error) {

	allMedia, err := cfg.db.GetAllMedia(context.Background())
	if err != nil {
		fmt.Println("Couldnt get media data from database")
	}
	trueData, err := cfg.pageDataArranger(allMedia, user)
	if err != nil {
		fmt.Println("Couldnt get truedata from pageDataArranger")
	}

	return trueData, nil

}

// Arranges a media slice from SQL queries into PageData structs
func (cfg *apiconfig) pageDataArranger(allMedia []database.Medium, user database.User) (PageData, error) {
	trueData := PageData{}

	for _, datapoint := range allMedia {

		trueData.User.Username = user.Username

		if datapoint.MediaType == "video" {

			videoPath := strings.TrimPrefix(datapoint.FilePath, "static")
			encodedPath := url.PathEscape(videoPath)

			trueData.Videos = append(trueData.Videos, MediaItem{Title: datapoint.MediaName, FilePath: encodedPath, Format: strings.TrimPrefix(datapoint.Format, ".")})

		}
		if datapoint.MediaType == "image" {

			imagePath := strings.TrimPrefix(datapoint.FilePath, "static")
			encodedPath := url.PathEscape(imagePath)
			trueData.Images = append(trueData.Images, MediaItem{Title: datapoint.MediaName, FilePath: encodedPath})

		}

	}
	return trueData, nil

}

// injects the searched dataset ONLY into the html, example: searched for "abcd", only injects those files that filenames have "abc"
func (cfg *apiconfig) searchedTemplateInjector(w http.ResponseWriter, r *http.Request, user database.User, trueData PageData) {
	// Parse the HTML templates
	tmpl, err := template.ParseFiles("index.html", "static/imageTemplate.html", "static/videoTemplate.html", "static/userDetailTemplate.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	// Execute the template with the data
	err = tmpl.Execute(w, trueData)
	if err != nil {
		http.Error(w, "Error rendering data template ddd", http.StatusInternalServerError)
		fmt.Println(err)
	}

}
