package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func InitConnection() (*pgx.Conn, error) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	bgCtx := context.Background()

	conn, err := pgx.Connect(bgCtx, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %w", err)
	}

	tempCtx, cancel := context.WithTimeout(bgCtx, 5*time.Second)
	defer cancel()

	if err := conn.Ping(tempCtx); err != nil {
		conn.Close(bgCtx)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return conn, nil
}
