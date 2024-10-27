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
	Title       string
	FilePath    string
	Format      string
	ID          int
	IsFavourite bool
}
type userInfo struct {
	Username string
}
type PageData struct {
	User   userInfo
	Title  string
	Videos []MediaItem
	Images []MediaItem
	Audios []MediaItem
}

// Dynamic Injection of Data function.
func (cfg *apiconfig) templateInjector(w http.ResponseWriter, r *http.Request, user database.User) {
	// Parse the HTML templates
	tmpl, err := template.ParseFiles("index.html", "static/imageTemplate.html", "static/videoTemplate.html", "static/userDetailTemplate.html", "static/audioTemplate.html")
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
	trimmedText := ""

	for _, datapoint := range allMedia {

		trueData.User.Username = user.Username

		if datapoint.MediaType == "video" {
			videoPath := strings.TrimPrefix(datapoint.FilePath, "static")
			encodedPath := url.PathEscape(videoPath)
			favourite, err := cfg.favouriteChecker(user, datapoint)
			if err != nil {
				fmt.Println("Couldnt get favourites", err)
			}
			if len(datapoint.MediaName) > 40 {
				trimmedText = datapoint.MediaName[:40]
			} else {
				trimmedText = datapoint.MediaName
			}

			trueData.Videos = append(trueData.Videos, MediaItem{
				Title:       trimmedText,
				FilePath:    encodedPath,
				Format:      strings.TrimPrefix(datapoint.Format, "."),
				ID:          int(datapoint.ID),
				IsFavourite: favourite})
		}
		if datapoint.MediaType == "image" {
			favourite, err := cfg.favouriteChecker(user, datapoint)
			if err != nil {
				fmt.Println("Couldnt get favourites", err)
			}
			if len(datapoint.MediaName) > 40 {
				trimmedText = datapoint.MediaName[:40]
			} else {
				trimmedText = datapoint.MediaName
			}
			imagePath := strings.TrimPrefix(datapoint.FilePath, "static")
			encodedPath := url.PathEscape(imagePath)
			trueData.Images = append(trueData.Images, MediaItem{
				Title:       trimmedText,
				FilePath:    encodedPath,
				ID:          int(datapoint.ID),
				IsFavourite: favourite})

		}
		if datapoint.MediaType == "audio" {
			favourite, err := cfg.favouriteChecker(user, datapoint)
			if err != nil {
				fmt.Println("Couldnt get favourites", err)
			}
			if len(datapoint.MediaName) > 40 {
				trimmedText = datapoint.MediaName[:40]
			} else {
				trimmedText = datapoint.MediaName
			}

			audioPath := strings.TrimPrefix(datapoint.FilePath, "static")
			encodedPath := url.PathEscape(audioPath)

			trueData.Audios = append(trueData.Audios, MediaItem{
				Title:       trimmedText,
				FilePath:    encodedPath,
				Format:      strings.TrimPrefix(datapoint.Format, "."),
				ID:          int(datapoint.ID),
				IsFavourite: favourite})

		}

	}

	return trueData, nil

}

// injects the searched dataset ONLY into the html, example: searched for "abcd", only injects those files that filenames have "abc"
func (cfg *apiconfig) searchedTemplateInjector(w http.ResponseWriter, r *http.Request, user database.User, trueData PageData) {
	// Parse the HTML templates
	tmpl, err := template.ParseFiles("index.html", "static/imageTemplate.html", "static/videoTemplate.html", "static/userDetailTemplate.html", "static/audioTemplate.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Execute the template with the data
	err = tmpl.Execute(w, trueData)
	if err != nil {
		http.Error(w, "Error rendering data template", http.StatusInternalServerError)
		fmt.Println(err)
	}

}

func (cfg *apiconfig) favouriteChecker(user database.User, datapoint database.Medium) (bool, error) {
	favourite := false
	// SQL query to check if following media
	d, err := cfg.db.GetFavouritedMedia(context.Background(), database.GetFavouritedMediaParams{UserID: user.ID, MediaID: datapoint.ID})
	if err != nil {
		fmt.Println("Error receiving favourited media")
		return favourite, err
	}
	if d == 1 {
		favourite = true

	}
	return favourite, nil

}

func (cfg *apiconfig) favouriteServer(user database.User) PageData {
	d, err := cfg.db.GetAllFavouriteMedia(context.Background(), user.ID)
	if err != nil {
		fmt.Println("Error couldnt get all favourite media", err)
	}
	trueData, err := cfg.pageDataArranger(d, user)
	if err != nil {
		fmt.Println("Didnt manage to arrange favourite media in pageDArranger", err)
	}
	return trueData

}
