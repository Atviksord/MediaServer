package main

import (
	"fmt"
	"net/http"

	"github.com/Atviksord/MediaServer/internal/database"
)

func (cfg *apiconfig) togglefavourite(w http.ResponseWriter, r *http.Request, user database.User) {
	fmt.Println("GOES INTO TOGGLE MODE")
	if r.Method == "POST" {
		fmt.Println("LIKED")
		mediaID := r.FormValue("mediaID")
		favourite := r.FormValue("favourite")
		fmt.Println(mediaID)
		fmt.Println(favourite)
		fmt.Println(user.Username)

	}

	return
}
