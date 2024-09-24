package main

import (
	"database/sql"
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
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in the environment variables")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	// Ping the database to confirm the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	// initialize database queries(SQL)
	dbQueries := database.New(db)
	cfg := &apiconfig{
		db: dbQueries,
	}

	PORT := os.Getenv("PORT")
	IP := os.Getenv("IP")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "0.0.0.0:" + PORT,
		Handler: mux,
	}
	cfg.handlerRegistry(mux)
	log.Printf("Server is starting on port %s\n", IP+":"+PORT)
	go directoryWatcherWorker("/static/Media")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
