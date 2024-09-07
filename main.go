package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Atviksord/MediaServer/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiconfig struct {
	db *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Couldn load env variables")
	}
	cfg := &apiconfig{}
	PORT := os.Getenv("PORT")
	IP := os.Getenv("IP")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "0.0.0.0:" + PORT,
		Handler: mux,
	}
	cfg.handlerRegistry(mux)
	log.Printf("Server is starting on port %s\n", IP+":"+PORT)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
