package services

import (
	"context"
	"fmt"
	"share-notes-app/internal/repositories"
	"time"

	"github.com/sirupsen/logrus"
)

type AuthorizationService interface {
	VerifyEmail(ctx context.Context, token string) (string, error)
}

type authorizationService struct {
	repo repositories.AuthorizationRepositories 
}

func NewAuthorizationsService(repo repositories.AuthorizationRepositories) AuthorizationService {
	return &authorizationService{repo:repo}
}

func (s *authorizationService) VerifyEmail(ctx context.Context, token string) (string, error) {
	// cek token kalau ada
	logrus.Info(token, "ini token dari controller")

	emailVerify, err := s.repo.GetToken(ctx, token)
	logrus.Info(emailVerify, "ini model dari repository")
	if err != nil {
		return "", fmt.Errorf("gagal mengambil token: %w", err)
	}

	if emailVerify == nil {
		return "", fmt.Errorf("token tidak ditemukan")
	}

	if emailVerify.IsUsed {
		return "", fmt.Errorf("token sudah digunakan")
	}

	if time.Now().After(emailVerify.ExpiresAt) {
		return "", fmt.Errorf("token sudah kadaluarsa")
	}


	// update status verifikasi
	err = s.repo.UpdateOneUsers(ctx, emailVerify)
	if err != nil {
		return "", fmt.Errorf("gagal mengupdate status: %w", err)
	} 

	return "berhasil meng-update status verifikasi", nil
}