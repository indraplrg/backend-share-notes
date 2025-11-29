package routes

import (
	"backend/internal/controllers"
	"backend/internal/repository"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuhtenticationRoute(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewAuthRepository(db)
	service := services.NewAuthService(repo)
	controller := controllers.NewAuthController(service)


	group := r.Group("/api/auth")
	{
		group.POST("/register", controller.Register)
		group.POST("/login", controller.Login)
		// group.POST("/verify", controllers.Register)
		// group.POST("/profile", controllers.Register)
	}

}