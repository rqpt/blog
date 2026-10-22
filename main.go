package main

import (
	"context"
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
	conn, err := db.InitConnection()
	if err != nil {
		log.Fatalf("Unable to establish database connection: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Serve the site.
	postHandler := handlers.NewPostHandler(conn)

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
			fmt.Sprintf(":%v", appPort),
			router,
		),
	)
}
