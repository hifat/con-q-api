package authRepo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
	"github.com/hifat/con-q-api/internal/app/model"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) authDomain.IAuthRepo {
	return &authRepo{db}
}

func (r *authRepo) Exists(ctx context.Context, authID uuid.UUID) (bool, error) {
	var exists bool
	return exists, r.db.WithContext(ctx).
		Model(&model.Auth{}).
		Select("COUNT(*) > 0").
		Where("id = ?", authID).
		Find(&exists).Error
}

func (r *authRepo) Register(ctx context.Context, req authDomain.ReqRegister) error {
	return r.db.WithContext(ctx).
		Create(&model.User{
			Username: req.Username,
			Password: req.Password,
			Name:     req.Name,
		}).Error
}

func (r *authRepo) Count(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64
	return count, r.db.WithContext(ctx).
		Model(&model.Auth{}).
		Where("user_id", userID).
		Count(&count).Error
}

func (r *authRepo) Save(ctx context.Context, req authDomain.ReqAuth) error {
	return r.db.WithContext(ctx).
		Save(&model.Auth{
			ID:        req.ID,
			Agent:     req.Agent,
			ClientIP:  req.ClientIP,
			ExpiresAt: req.ExpiresAt,

			UserID: req.UserID,
		}).Error
}

func (r *authRepo) Delete(ctx context.Context, authID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", authID).
		Delete(&model.Auth{}).Error
}

func (r *authRepo) RemoveTokenExpires(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Where("expires_at <= ?", time.Now().Format(time.RFC3339)).
		Delete(&model.Auth{}).Error
}
