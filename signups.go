package main

// Ideas for improvement TODO: Signups also hash passwords before writing to DB and retrieval functions will unhash to match
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Atviksord/MediaServer/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiconfig) signupHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body %v", err)
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		err = r.ParseForm()
		if err != nil {
			fmt.Printf("Error parsing form %v", err)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")

		// hash the password
		truePass := cfg.passwordHasher(password)

		// Check DB for match IF NOT EXIST create user in DB and automatically LOG IN (jwt creation, also hash password etc)
		_, err = cfg.db.GetUser(r.Context(), username)
		if err != nil {
			cfg.db.CreateUser(r.Context(), database.CreateUserParams{
				Username:  username,
				Password:  string(truePass),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC()})

			fmt.Printf("User does not exist, creating new user %s with the password %s", username, password)
			http.ServeFile(w, r, "./static/login.html")

			return
		}

	}

}

func (cfg *apiconfig) passwordHasher(password string) []byte {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println("Unable to hash pass")
	}
	return encryptedPass
}
