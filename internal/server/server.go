package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func NewServer(router http.Handler) *http.Server {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT is not set")
	}

	server := http.Server{
		Addr:    fmt.Sprint(":", appPort),
		Handler: router,
	}

	return &server
}
