package repository

import (
	"context"
	"errors"

	"lapakita-backend/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Preload("IdentityProfile").
		First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) UpdatePassword(ctx context.Context, userID uuid.UUID, newPasswordHash string) error {
	return r.db.WithContext(ctx).
		Model(&entity.User{}).
		Where("id = ?", userID).
		Update("password_hash", newPasswordHash).Error
}

func (r *UserRepository) FindIdentityByNIK(ctx context.Context, nik string) (*entity.UserIdentityProfile, error) {
	var profile entity.UserIdentityProfile
	err := r.db.WithContext(ctx).First(&profile, "nik = ?", nik).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *UserRepository) UpsertIdentityProfile(ctx context.Context, profile *entity.UserIdentityProfile) error {
	var existing entity.UserIdentityProfile
	err := r.db.WithContext(ctx).First(&existing, "user_id = ?", profile.UserID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return r.db.WithContext(ctx).Create(profile).Error
		}
		return err
	}

	profile.ID = existing.ID
	return r.db.WithContext(ctx).Model(&existing).Updates(map[string]interface{}{
		"full_name_ktp": profile.FullNameKTP,
		"nik":           profile.NIK,
		"ktp_photo_url": profile.KTPPhotoURL,
		"domicile_city": profile.DomicileCity,
	}).Error
}

func (r *UserRepository) DeleteIdentityProfile(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&entity.UserIdentityProfile{}).Error
}

func (r *UserRepository) GetDocument(ctx context.Context, userID uuid.UUID, name string, nik string, page int, limit int) ([]entity.UserIdentityProfile, int64, error) {
	var documents []entity.UserIdentityProfile
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.UserIdentityProfile{}).Where("user_id = ?", userID)

	if name != "" {
		query = query.Where("full_name_ktp ILIKE ?", "%"+name+"%")
	}
	if nik != "" {
		query = query.Where("nik ILIKE ?", "%"+nik+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&documents).Error
	if err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}
