package repository

import (
	"context"
	"errors"
	"math"
	"sort"
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

	var indexedList []dto.PhoneNumberItem
	var selectedItem *dto.PhoneNumberItem

	// 1. Filter Search & Map Index Asli DB
	for dbIdx, p := range user.PhoneNumbers {
		if req.Number != "" && !strings.Contains(strings.ToLower(p.Number), strings.ToLower(req.Number)) {
			continue
		}

		item := dto.PhoneNumberItem{
			Index:     dbIdx,
			Number:    p.Number,
			IsPrimary: p.IsPrimary,
			Roles:     p.Roles,
		}

		// SelectedID berupa string nomor HP (e.g., "+628123456789" atau "08123456789")
		if req.SelectedID != "" && p.Number == req.SelectedID {
			itemCopy := item
			selectedItem = &itemCopy
			continue
		}

		indexedList = append(indexedList, item)
	}

	// 2. Sort Primary First
	sort.SliceStable(indexedList, func(i, j int) bool {
		if indexedList[i].IsPrimary {
			return true
		}
		if indexedList[j].IsPrimary {
			return false
		}
		return false
	})

	// 3. Inject Selected Item ke Posisi Terdepan (Page 1 Anchor)
	if selectedItem != nil {
		indexedList = append([]dto.PhoneNumberItem{*selectedItem}, indexedList...)
	}

	totalItems := len(indexedList)
	if totalItems == 0 {
		meta := api.PaginationMeta{
			TotalItems:  0,
			TotalPages:  0,
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: false,
		}
		return []dto.PhoneNumberItem{}, meta, nil
	}

	// 4. In-Memory Slicing (Direction UP/DOWN handling)
	var startIndex, endIndex int

	if req.Page > 1 {
		startIndex = (req.Page - 1) * req.Limit
		endIndex = startIndex + req.Limit
	} else if req.Page < 1 {
		pageOffset := (-req.Page) * req.Limit
		startIndex = pageOffset - req.Limit
		endIndex = pageOffset
	} else {
		startIndex = 0
		endIndex = req.Limit
	}

	if startIndex < 0 {
		startIndex = 0
	}
	if startIndex >= totalItems {
		meta := api.PaginationMeta{
			TotalItems:  totalItems,
			TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: req.Page > 1 || req.Page < 0,
		}
		return []dto.PhoneNumberItem{}, meta, nil
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	pagedItems := indexedList[startIndex:endIndex]
	totalPages := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

	meta := api.PaginationMeta{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: endIndex < totalItems,
		HasPrevPage: req.Page > 1 || req.Page < 0,
	}

	return pagedItems, meta, nil
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
