package repositories

import (
	"context"
	"errors"
	"share-notes-app/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthorizationRepositories interface {
	GetToken(ctx context.Context, token string) (*models.EmailVerification ,error)
	UpdateOneUsers(ctx context.Context, emailVerify *models.EmailVerification) error
}

type authorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) AuthorizationRepositories {
	return &authorizationRepository{db:db}
}

func (r *authorizationRepository) GetToken(ctx context.Context, token string) (*models.EmailVerification ,error)  {
	var EmailVerification models.EmailVerification 
	logrus.Info(token, "ini token dari service")

	if err := r.db.WithContext(ctx).Where("token = ?", token).First(&EmailVerification).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &EmailVerification, nil
}

func (r *authorizationRepository) UpdateOneUsers(ctx context.Context, emailVerify *models.EmailVerification) error {
	logrus.Info(emailVerify.UserId, "ini id dari tabel email")
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// update user
		if err := tx.Model(&models.User{}).Where("id = ?", emailVerify.UserId).Update("is_verified", true).Error; err != nil {
			return err
		}

		// update token
		if err := tx.Model(&models.EmailVerification{}).Where("id = ?", emailVerify.ID).Update("is_used", true).Error; err != nil {
			return err
		}

		return nil
	})
}