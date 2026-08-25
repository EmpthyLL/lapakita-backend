package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"lapakita-backend/internal/entity"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepository interface {
	// Database Operations
	CreateUser(ctx context.Context, user *entity.User) error
	FindUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error

	// Redis OTP Operations
	SetOTP(ctx context.Context, mode string, email string, code string, ttl time.Duration) error
	GetOTP(ctx context.Context, mode string, email string) (string, error)
	DeleteOTP(ctx context.Context, mode string, email string) error

	// Redis Verification Token Operations
	SetVerificationToken(ctx context.Context, mode string, email string, token string, ttl time.Duration) error
	GetVerificationToken(ctx context.Context, mode string, email string) (string, error)
	DeleteVerificationToken(ctx context.Context, mode string, email string) error
}

type authRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewAuthRepository(db *gorm.DB, rdb *redis.Client) AuthRepository {
	return &authRepository{
		db:  db,
		rdb: rdb,
	}
}

func (r *authRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Preload("IdentityProfile").
		Preload("BankAccounts").
		First(&user, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *authRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Preload("IdentityProfile").
		Preload("BankAccounts").
		First(&user, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *authRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// -----------------------------------------------------------------------------
// REDIS OTP & VERIFICATION HELPERS
// -----------------------------------------------------------------------------

func (r *authRepository) SetOTP(ctx context.Context, mode string, email string, code string, ttl time.Duration) error {
	key := fmt.Sprintf("otp:%s:%s", mode, email)
	return r.rdb.Set(ctx, key, code, ttl).Err()
}

func (r *authRepository) GetOTP(ctx context.Context, mode string, email string) (string, error) {
	key := fmt.Sprintf("otp:%s:%s", mode, email)
	return r.rdb.Get(ctx, key).Result()
}

func (r *authRepository) DeleteOTP(ctx context.Context, mode string, email string) error {
	key := fmt.Sprintf("otp:%s:%s", mode, email)
	return r.rdb.Del(ctx, key).Err()
}

func (r *authRepository) SetVerificationToken(ctx context.Context, mode string, email string, token string, ttl time.Duration) error {
	key := fmt.Sprintf("verified:%s:%s", mode, email)
	return r.rdb.Set(ctx, key, token, ttl).Err()
}

func (r *authRepository) GetVerificationToken(ctx context.Context, mode string, email string) (string, error) {
	key := fmt.Sprintf("verified:%s:%s", mode, email)
	return r.rdb.Get(ctx, key).Result()
}

func (r *authRepository) DeleteVerificationToken(ctx context.Context, mode string, email string) error {
	key := fmt.Sprintf("verified:%s:%s", mode, email)
	return r.rdb.Del(ctx, key).Err()
}
