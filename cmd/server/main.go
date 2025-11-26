package main

import (
	"backend/internal/database"
	"log"
)

func main() {
	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}