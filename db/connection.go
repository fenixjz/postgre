package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"postgre/config"
)

var Pool *pgxpool.Pool

func Connect() {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	Pool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
