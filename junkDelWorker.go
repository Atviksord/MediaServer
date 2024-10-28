package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func (cfg *apiconfig) junkDeleter(filepathz string) {
	for {
		// Go Worker routine to delete useless file in media folder, Zone identifier files etc.
		time.Sleep(5 * time.Minute)

		allfiles, err := os.ReadDir(filepathz)
		if err != nil {
			fmt.Println("Error reading directory from", filepathz)
		}
		for _, singlefile := range allfiles {
			if !singlefile.IsDir() {
				detailFile, err := singlefile.Info()
				if err != nil {
					fmt.Println("Unable to get info on file", err)
				}

				filename := detailFile.Name()
				extension := filepath.Ext(filename)
				if extension == ".Identifier" {
					err := os.Remove(filepath.Join(filepathz, filename))
					if err != nil {
						fmt.Println("Failed to remove file", err)
					}
				}

			}

		}
	}
}
