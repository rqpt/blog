package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rqpt/blog/internal/db"
)

func main() {
	// Establish a connection to the database.
	conn, err := db.InitConnection()
	if err != nil {
		log.Fatalf("Unable to establish database connection: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Serve a basic homepage.
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<body>Welcome!</body>")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", appPort), nil))
}
