package services

import (
	"context"
	"errors"
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
	emailVerify, err := s.repo.GetToken(ctx, token)

	if err != nil {
		logrus.WithError(err)
		return "", errors.New("gagal mengambil token")
	}

	if emailVerify == nil {
		return "", errors.New("token tidak ditemukan")
	}

	if emailVerify.IsUsed {
		return "", errors.New("token sudah digunakan")
	}

	if time.Now().After(emailVerify.ExpiresAt) {
		return "", errors.New("token sudah kadaluarsa")
	}


	// update status verifikasi
	err = s.repo.UpdateOneUsers(ctx, emailVerify)
	if err != nil {
		logrus.WithError(err)
		return "", errors.New("gagal mengupdate status")
	} 

	return "berhasil meng-update status verifikasi", nil
}