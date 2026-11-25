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

func (r *BusinessRepository) List(ctx context.Context, req dto.ListBusinessRequest) ([]entity.Business, map[uuid.UUID]int64, api.PaginationMeta, error) {
	var businesses []entity.Business
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).Model(&entity.Business{}).
		Preload("BusinessType").
		Preload("Fronts", func(db *gorm.DB) *gorm.DB {
			rankedSubquery := `
				SELECT * FROM (
					SELECT *,
						ROW_NUMBER() OVER (
							PARTITION BY business_id
							ORDER BY
								CASE platform
									WHEN 'offline_stall' THEN 1
									WHEN 'website' THEN 2
									ELSE 3
								END ASC,
								created_at ASC
						) AS rn
					FROM business_fronts
					WHERE deleted_at IS NULL
				) ranked
				WHERE ranked.rn <= 5
			`
			return db.Table("(" + rankedSubquery + ") AS business_fronts").
				Order("CASE platform WHEN 'offline_stall' THEN 1 WHEN 'website' THEN 2 ELSE 3 END ASC, created_at ASC")
		}).
		Where("deleted_at IS NULL AND user_id = ?", req.UserID)

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
		return nil, nil, api.PaginationMeta{}, err
	}

	// Hitung total fronts per bisnis dalam 1 query agar frontend tahu jumlah aslinya
	frontCounts := make(map[uuid.UUID]int64, len(businesses))
	if len(businesses) > 0 {
		businessIDs := make([]uuid.UUID, 0, len(businesses))
		for _, b := range businesses {
			businessIDs = append(businessIDs, b.ID)
		}

		type countResult struct {
			BusinessID uuid.UUID
			Total      int64
		}
		var results []countResult
		_ = r.db.WithContext(ctx).
			Model(&entity.BusinessFront{}).
			Select("business_id, COUNT(*) AS total").
			Where("business_id IN ? AND deleted_at IS NULL", businessIDs).
			Group("business_id").
			Scan(&results).Error

		for _, res := range results {
			frontCounts[res.BusinessID] = res.Total
		}
	}

	return businesses, frontCounts, meta, nil
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

// =============================================================================
// BUSINESS FRONTS REPOSITORY METHODS
// =============================================================================

func (r *BusinessRepository) CreateFront(ctx context.Context, front *entity.BusinessFront) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Create(front).Error
	})
}

func (r *BusinessRepository) FindFrontsByBusinessID(
	ctx context.Context,
	businessID uuid.UUID,
	req dto.GetBusinessFrontsRequest,
) ([]entity.BusinessFront, api.PaginationMeta, error) {
	var fronts []entity.BusinessFront
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).Model(&entity.BusinessFront{}).
		Where("business_id = ? AND deleted_at IS NULL", businessID)

	if platform := strings.TrimSpace(req.Platform); platform != "" {
		query = query.Where("platform = ?", platform)
	}

	if search := strings.TrimSpace(req.Search); search != "" {
		query = query.Where("LOWER(store_name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"business_fronts",
		&fronts,
		&meta,
		"CASE platform WHEN 'offline_stall' THEN 1 WHEN 'website' THEN 2 ELSE 3 END ASC, created_at ASC",
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return fronts, meta, nil
}

func (r *BusinessRepository) FindFrontByID(ctx context.Context, frontID uuid.UUID) (*entity.BusinessFront, error) {
	var front entity.BusinessFront
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", frontID).
		First(&front).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &front, err
}

func (r *BusinessRepository) UpdateFront(ctx context.Context, front *entity.BusinessFront) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Save(front).Error
	})
}

func (r *BusinessRepository) DeleteFront(ctx context.Context, frontID uuid.UUID) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Model(&entity.BusinessFront{}).
			Where("id = ?", frontID).
			Update("deleted_at", gorm.Expr("NOW()")).Error
	})
}
