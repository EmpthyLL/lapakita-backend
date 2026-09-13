package repository

import (
	"context"
	"errors"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"

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

// Check jika nomor dokumen ini sudah dipakai di manapun (Global Check)
func (r *UserRepository) FindIdentityByDocumentNumber(ctx context.Context, docNumber string) (*entity.UserIdentityProfile, error) {
	var profile entity.UserIdentityProfile
	err := r.db.WithContext(ctx).First(&profile, "document_number = ?", docNumber).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

// Check khusus apakah user spesifik ini sudah pernah menambahkan document_number tersebut
func (r *UserRepository) FindIdentityByUserIDAndDocNumber(ctx context.Context, userID uuid.UUID, docNumber string) (*entity.UserIdentityProfile, error) {
	var profile entity.UserIdentityProfile
	err := r.db.WithContext(ctx).First(&profile, "user_id = ? AND document_number = ?", userID, docNumber).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *UserRepository) CreateIdentityProfile(ctx context.Context, profile *entity.UserIdentityProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *UserRepository) DeleteIdentityProfile(ctx context.Context, userID uuid.UUID, documentID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ? AND user_id = ?", documentID, userID).Delete(&entity.UserIdentityProfile{}).Error
}

func (r *UserRepository) GetDocument(ctx context.Context, userID uuid.UUID, req dto.GetDocumentRequest) ([]entity.UserIdentityProfile, int64, error) {
	var documents []entity.UserIdentityProfile
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.UserIdentityProfile{}).Where("user_id = ?", userID)

	if req.Name != "" {
		query = query.Where("full_name_identity ILIKE ?", "%"+req.Name+"%")
	}
	if req.DocumentNumber != "" {
		query = query.Where("document_number ILIKE ?", "%"+req.DocumentNumber+"%")
	}
	if req.DocumentType != "" {
		query = query.Where("document_type = ?", req.DocumentType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.Limit
	err := query.Offset(offset).Limit(req.Limit).Order("created_at DESC").Find(&documents).Error
	if err != nil {
		return nil, 0, err
	}

	return documents, total, nil
}
