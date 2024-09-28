package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Directory watcher to watch changes in local files to upload to DB
func directoryWatcherWorker(dirPath string) {
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

			// Get file info
			fileName := filepath.Base(event.Name)
			filePath := event.Name
			fileExt := filepath.Ext(event.Name)
			fileType := "unknown"

			// Determine the file type based on extension
			switch strings.ToLower(fileExt) {
			case ".jpg", ".jpeg", ".png", ".gif":
				fileType = "image"
			case ".mp4", ".avi", ".mov":
				fileType = "video"
			case ".mp3", ".wav":
				fileType = "audio"

			}

			// Handle different types of events
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Printf("File created: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				fmt.Printf("File deleted: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("File modified: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
			}
			if event.Op&fsnotify.Rename == fsnotify.Rename {
				fmt.Printf("File renamed: Name: %s, Path: %s, Type: %s, Format: %s at %s\n", fileName, filePath, fileType, fileExt, time.Now())
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}
