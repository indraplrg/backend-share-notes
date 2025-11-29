package models

import "github.com/golang-jwt/jwt/v5"

type Role string

const (
	AdminRole Role = "admin"
	UserRole Role = "user"		
)

type Payload struct {
	ID string `json:"id"`
	Role Role `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type ResponseToken struct {
	AccessToken string
	RefreshToken string
}