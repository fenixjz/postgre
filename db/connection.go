package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"postgre/config"
)

var Conn *pgx.Conn

// Connect establishes connection to the PostgreSQL database.
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
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

// Close safely closes the database connection.
func Close() {
	if Conn != nil {
		if err := Conn.Close(context.Background()); err != nil {
			log.Printf("Error closing DB connection: %v", err)
		}
	}
}
