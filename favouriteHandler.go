package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Atviksord/MediaServer/internal/database"
)

func (cfg *apiconfig) togglefavourite(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println("GOES INTO TOGGLE MODE")
	if r.Method == "POST" {
		fmt.Println("LIKED")
		StringMediaID := r.FormValue("mediaID")
		favourite := r.FormValue("favourite")
		fmt.Println(StringMediaID)
		mediaID, err := strconv.Atoi(StringMediaID)
		if err != nil {
			fmt.Println("Failed to convert stringMediaID into MediaIT (integer)")
		}
		fmt.Println(favourite)
		fmt.Println(user.Username)
		if favourite == "false" {
			fmt.Println("Generating a new favourite")
			_, err := cfg.db.AddFavourite(r.Context(), database.AddFavouriteParams{UserID: user.ID, MediaID: int32(mediaID)})
			if err != nil {
				fmt.Println("Unable to add favourite", err)
			}

		} else {
			fmt.Println("Removing a favourite")
			_, err := cfg.db.DeleteFavourite(r.Context(), database.DeleteFavouriteParams{UserID: user.ID, MediaID: int32(mediaID)})
			if err != nil {
				fmt.Println("Unable to remove favourite", err)
			}
		}
		// RESERVE POINT
		cfg.templateInjector(w, r, user)
		// Need to fix bug, make sure the unfilled heart disappears on toggle.

	}
}
