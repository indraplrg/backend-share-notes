package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, email string, password string) (*models.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
    return &authService{repo: repo}
}

func (s *authService) Register(ctx context.Context, email string, password string) (*models.User, error) {
	
	// cek email terdaftar?
	existing, _ := s.repo.FindByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email ini sudah terdaftar")
	}

	// hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// create user
	user := &models.User{
		Email: email,
		Password: string(hashed),
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("gagal membuat akun")
	}

	return user, nil 
}