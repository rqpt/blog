package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rqpt/blog/internal/db"
	"github.com/rqpt/blog/internal/handlers"
	"github.com/rqpt/blog/internal/router"
)

func main() {
	// Establish a connection to the database.
	dbPool, err := db.Connect()
	if err != nil {
		log.Fatalf("Unable to establish database connection: %v\n", err)
	}
	defer dbPool.Close()

	// Serve the site.
	postHandler := handlers.NewPostHandler(dbPool)

	router := router.NewRouter(
		router.Handlers{
			Posts: postHandler,
		},
	)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set")
	}

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprint(":", appPort),
			router,
		),
	)
}
