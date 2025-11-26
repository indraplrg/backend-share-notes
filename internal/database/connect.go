package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := "user=void dbname=share_notes sslmode=verify-full"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Failed to open Database connection: %v", err)
	}
	
	return db, nil
}