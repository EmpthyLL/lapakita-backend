package repository

import (
	"context"
	"errors"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) *BusinessRepository {
	return &BusinessRepository{db: db}
}

func (r *BusinessRepository) Create(ctx context.Context, business *entity.Business) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Create(business).Error
	})
}

func (r *BusinessRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Business, error) {
	var business entity.Business
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").First(&business, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &business, err
}

func (r *BusinessRepository) FindByIDWithPreload(ctx context.Context, id uuid.UUID) (*entity.Business, error) {
	var business entity.Business
	err := r.db.WithContext(ctx).
		Preload("User").
		Preload("BusinessType").
		Where("deleted_at IS NULL").
		First(&business, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &business, err
}

func (r *BusinessRepository) List(ctx context.Context, req dto.ListBusinessRequest) ([]entity.Business, api.PaginationMeta, error) {
	var businesses []entity.Business
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).Model(&entity.Business{}).
		Preload("BusinessType").
		Where("deleted_at IS NULL")

	if req.UserID != nil {
		query = query.Where("user_id = ?", *req.UserID)
	}

	if req.BusinessTypeID != nil {
		query = query.Where("business_type_id = ?", *req.BusinessTypeID)
	}

	if search := strings.TrimSpace(req.Search); search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"businesses",
		&businesses,
		&meta,
		"created_at DESC",
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return businesses, meta, nil
}

func (r *BusinessRepository) Update(ctx context.Context, business *entity.Business) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Save(business).Error
	})
}

func (r *BusinessRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Model(&entity.Business{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
	})
}
