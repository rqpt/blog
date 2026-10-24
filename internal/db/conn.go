package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitConnection() (*pgxpool.Pool, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	bgCtx := context.Background()

	pool, err := pgxpool.New(bgCtx, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	tempCtx, cancel := context.WithTimeout(bgCtx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(tempCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}
