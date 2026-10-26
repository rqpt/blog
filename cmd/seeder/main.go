package main

import (
	"context"
	"log"

	"github.com/rqpt/blog/internal/db"
)

func main() {
	ctx := context.Background()

	dbPool, err := db.Connect()
	if err != nil {
		log.Fatalf("Unable to establish database connection: %v\n", err)
	}
	defer dbPool.Close()

	if err := db.Seed(ctx, dbPool); err != nil {
		log.Fatal(err)
	}
}
