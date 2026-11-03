package main

import (
	"log"

	"github.com/rqpt/blog/internal/db"
	"github.com/rqpt/blog/internal/handlers"
	"github.com/rqpt/blog/internal/router"
	"github.com/rqpt/blog/internal/server"
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

	server := server.New(router)

	log.Fatal(server.ListenAndServe())
}
