package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	GenerateToken(user *models.User) (*models.ResponseToken, error)
	Register(ctx context.Context, email string, password string) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.ResponseToken, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
    return &authService{repo: repo}
}

func (s *authService) GenerateToken(user *models.User) (*models.ResponseToken, error) {
	accessTokenPayload := models.Payload{
		ID: user.ID.String(),
		Username: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	
	refreshTokenPayload := models.Payload{
		ID: user.ID.String(),
		Username: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10080)),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenPayload)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenPayload)

	signedAccessToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, errors.New("gagal membuat Akses token")
	}

	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return nil, errors.New("gagal membuat Refresh token")
	}

	token := &models.ResponseToken{
		AccessToken: signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	return token, nil
}

func (s *authService) Register(ctx context.Context, email string, password string) (*models.User, error) {
	
	// cek email terdaftar?
	existing, err := s.repo.FindByEmail(ctx, email)
	

	if err != nil {
		return nil, fmt.Errorf("gagal mengecek user: %w", err)
	}
	
	if existing != nil {
		return nil, errors.New("email ini sudah terdaftar")
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal meng-hash password")
	}


	// create user
	user := &models.User{
		Email: email,
		Password: string(hashed),
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.New("gagal membuat akun")
	}

	return user, nil 
}

func (s *authService) Login(ctx context.Context, email string, password string) (*models.ResponseToken, error) {

	// cek email terdaftar?
	existing, err := s.repo.FindByEmail(ctx, email)
	
	if err != nil {
		return nil, fmt.Errorf("gagal mengecek user: %w", err)
	}

	if existing == nil {
		return nil, errors.New("email ini Belum terdaftar")
	}	

	// cek password benar?
	isPasswordTrue := bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(password))
	if isPasswordTrue != nil{
		return nil, errors.New("password yang anda masukkan salah")
	}

	// buat token
	token, err := s.GenerateToken(existing)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat token: %w", err)
	}

	return token, nil

}

