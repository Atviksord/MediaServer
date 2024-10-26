package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Atviksord/MediaServer/internal/database"
	"github.com/fsnotify/fsnotify"
)

// Directory watcher to watch changes in local files to upload to DB
func (cfg *apiconfig) directoryWatcherWorker(dirPath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Add the directory to the watcher
	err = watcher.Add(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// Infinite loop to monitor directory events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			// Get info about the file, filepath, name, type extension and file type(video,image, etc)
			fileName := filepath.Base(event.Name)
			filePath := event.Name
			fileExt := filepath.Ext(event.Name)
			fileType := "unknown"

			switch strings.ToLower(fileExt) {
			case ".jpg", ".jpeg", ".png", ".gif":
				fileType = "image"
			case ".mp4", ".avi", ".mov":
				fileType = "video"
			case ".mp3", ".wav", ".m4a":
				fileType = "audio"
			}

			if event.Op&fsnotify.Create == fsnotify.Create && fileType != "unknown" {
				fmt.Printf("File created: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
				_, err := cfg.db.AddMedia(context.Background(), database.AddMediaParams{
					MediaName:  fileName,
					MediaType:  fileType,
					FilePath:   filePath,
					Format:     fileExt,
					UploadDate: sql.NullTime{Time: time.Now().UTC(), Valid: true},
				})
				if err != nil {
					fmt.Println(err)
				}

			}
			// delete media FROM DB if it detects it has been removed
			if event.Op&fsnotify.Remove == fsnotify.Remove && fileType != "unknown" {
				fmt.Printf("File deleted: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
				_, err := cfg.db.DeleteMedia(context.Background(), filePath)
				fmt.Println(filePath)
				if err != nil {
					fmt.Println("Error removing from DB", err)
				}
			}

			if event.Op&fsnotify.Rename == fsnotify.Rename && fileType != "unknown" {
				// Check if the file still exists
				if _, err := os.Stat(event.Name); os.IsNotExist(err) {
					// Treat as delete if the file does not exist
					fmt.Printf("File deleted: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
					_, err := cfg.db.DeleteMedia(context.Background(), filePath)
					fmt.Println(filePath)
					if err != nil {
						fmt.Println("Error removing from DB", err)
					}

				} else {
					fmt.Printf("File renamed: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}
