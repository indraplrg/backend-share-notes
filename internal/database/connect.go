package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDBConnection() (*gorm.DB, error) {
	connStr := "user=void password=voidajalah port=5432 dbname=share_notes sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Failed to open Database connection: %v", err)
	}
	
	return db, nil
}