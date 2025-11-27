package main

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	r := gin.Default()

	
	routes.AuhtenticationRoute(r, db)

	r.Run(":3000")
}