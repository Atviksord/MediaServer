package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
)

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
			// Handle different types of events
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Printf("File %s was created at %s\n", event.Name, time.Now())
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				fmt.Printf("File %s was deleted at %s\n", event.Name, time.Now())
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("File %s was modified at %s\n", event.Name, time.Now())
			}
			if event.Op&fsnotify.Rename == fsnotify.Rename {
				fmt.Printf("File %s was renamed at %s\n", event.Name, time.Now())
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}
