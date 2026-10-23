package repository

import (
	"context"
	"errors"

	"lapakita-backend/internal/entity"
	"lapakita-backend/pkg/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
}

func (r *AuthRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Preload("IdentityProfile").
		Preload("PaymentMethod").
		First(&user, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *AuthRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Preload("IdentityProfile").
		Preload("PaymentMethod").
		First(&user, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *AuthRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Save(user).Error
	})
}
