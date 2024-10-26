package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Atviksord/MediaServer/internal/database"
)

// Is not properly toggling between false and true values from template
func (cfg *apiconfig) togglefavourite(w http.ResponseWriter, r *http.Request, user database.User) {
	if r.Method == "POST" {
		favourite := r.FormValue("favourite")
		StringMediaID := r.FormValue("mediaID")
		mediaID, err := strconv.Atoi(StringMediaID)

		if err != nil {
			fmt.Println("Failed to convert stringMediaID into MediaIT (integer)")
		}
		if favourite == "" {
			fmt.Println("Removing a favourite")
			_, err := cfg.db.DeleteFavourite(r.Context(), database.DeleteFavouriteParams{UserID: user.ID, MediaID: int32(mediaID)})
			if err != nil {
				fmt.Println("Unable to remove favourite", err)
			}
		}
		if favourite == "true" {
			fmt.Println("Generating a new favourite")
			_, err := cfg.db.AddFavourite(r.Context(), database.AddFavouriteParams{UserID: user.ID, MediaID: int32(mediaID)})
			if err != nil {
				fmt.Println("Unable to add favourite", err)
			}
		}
		// RESERVE POINT
		cfg.templateInjector(w, r, user)
		// Need to fix bug, make sure the unfilled heart disappears on toggle.

	}
}
