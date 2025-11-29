package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Email string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    Role Role `gorm:"type:varchar(20);default:'user';not null"`
    RefreshToken *string `gorm:"type:text"`
    CreatedAt time.Time
}

type UserResponse struct {
    ID uuid.UUID `json:"id"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type Response struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

type UserRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}