package main

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load config envinroment
	err := godotenv.Load() 
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	// koneksi ke database
	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	r := gin.Default()

	
	routes.AuhtenticationRoute(r, db)

	r.Run(":" + os.Getenv("PORT"))
}