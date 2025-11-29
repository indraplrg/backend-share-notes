package controllers

import (
	"backend/internal/models"
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req models.UserRequest
	
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	user, err := c.service.Register(ctx, req.Email, req.Password); 
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})

		return
	}

	response := models.UserResponse{
		ID: user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Berhasil membuat akun",
		Data: response,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.UserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	token, err := c.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Berhasil Login",
		Data: token,
	})
}
