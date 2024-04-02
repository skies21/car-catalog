package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	connStr := os.Getenv("DB_CONNECTION_STRING")
	if connStr == "" {
		log.Fatal("DB_CONNECTION_STRING is not set in .env file")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db, nil
}
