package repository

import (
	"context"
	"errors"
	"sort"
	"strconv"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

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

func (r *UserRepository) GetPhoneNumbers(ctx context.Context, userID uuid.UUID, req *dto.GetPhoneNumbersRequest) ([]dto.PhoneNumberItem, api.PaginationMeta, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Select("id, phone_numbers").
		First(&user, "id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api.PaginationMeta{}, errors.New("user not found")
		}
		return nil, api.PaginationMeta{}, err
	}

	req.SetDefaults()

	// 1. Map ke DTO & Filter Search (Bersih tanpa modifikasi array ganda)
	var filteredList []dto.PhoneNumberItem
	for dbIdx, p := range user.PhoneNumbers {
		phoneNumber := p.DialCode + p.Number
		if req.Search != "" && !strings.Contains(strings.ToLower(phoneNumber), strings.ToLower(req.Search)) {
			continue
		}

		filteredList = append(filteredList, dto.PhoneNumberItem{
			Index:     dbIdx,
			DialCode:  p.DialCode,
			Number:    p.Number,
			IsPrimary: p.IsPrimary,
			Roles:     p.Roles,
		})
	}

	// 2. Sort Visual: Primary selalu di atas
	sort.SliceStable(filteredList, func(i, j int) bool {
		if filteredList[i].IsPrimary {
			return true
		}
		if filteredList[j].IsPrimary {
			return false
		}
		return false
	})

	var resultSlice []dto.PhoneNumberItem
	var meta api.PaginationMeta

	// 3. Eksekusi AutoPaginateSlice (Pencocokan presisi string nomor HP)
	database.AutoPaginateSlice(
		filteredList,
		req.BasePaginationRequest,
		&resultSlice,
		&meta,
		func(item dto.PhoneNumberItem, selectedID string) bool {
			selectedIndex, err := strconv.Atoi(selectedID)
			return err == nil && item.Index == selectedIndex
		},
	)

	return resultSlice, meta, nil
}

func (r *UserRepository) CreateIdentityProfile(ctx context.Context, profile *entity.UserIdentityProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *UserRepository) DeleteIdentityProfile(ctx context.Context, userID uuid.UUID, documentID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ? AND user_id = ?", documentID, userID).Delete(&entity.UserIdentityProfile{}).Error
}

func (r *UserRepository) GetDocument(ctx context.Context, userID uuid.UUID, req *dto.GetDocumentRequest) ([]entity.UserIdentityProfile, api.PaginationMeta, error) {
	var documents []entity.UserIdentityProfile
	var meta api.PaginationMeta

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

	err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"user_identity_profiles",
		&documents,
		&meta,
		"created_at DESC",
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return documents, meta, nil
}
