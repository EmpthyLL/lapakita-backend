package repository

import (
	"context"
	"errors"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/wallet/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PayoutMethodRepository struct {
	db *gorm.DB
}

func NewPayoutMethodRepository(db *gorm.DB) *PayoutMethodRepository {
	return &PayoutMethodRepository{db: db}
}

func (r *PayoutMethodRepository) GetByUserID(ctx context.Context, userID uuid.UUID, req dto.GetPayoutMethodsRequest) ([]entity.PayoutMethod, api.PaginationMeta, error) {
	var methods []entity.PayoutMethod
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).
		Model(&entity.PayoutMethod{}).
		Where("user_id = ?", userID)
	if err := database.AutoPaginate(query, req.BasePaginationRequest, "payout_methods", &methods, &meta, ""); err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return methods, meta, nil
}

func (r *PayoutMethodRepository) Create(ctx context.Context, method *entity.PayoutMethod) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		var user entity.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&user, "id = ?", method.UserID).Error; err != nil {
			return err
		}

		var count int64
		if err := tx.Model(&entity.PayoutMethod{}).Where("user_id = ?", method.UserID).Count(&count).Error; err != nil {
			return err
		}
		method.IsPrimary = count == 0
		return tx.Create(method).Error
	})
}

func (r *PayoutMethodRepository) Delete(ctx context.Context, userID, methodID uuid.UUID) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		var user entity.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&user, "id = ?", userID).Error; err != nil {
			return err
		}
		var method entity.PayoutMethod
		if err := tx.Where("user_id = ? AND id = ?", userID, methodID).First(&method).Error; err != nil {
			return err
		}
		if err := tx.Delete(&method).Error; err != nil {
			return err
		}
		if !method.IsPrimary {
			return nil
		}
		if err := tx.Model(&entity.PayoutMethod{}).
			Where("user_id = ?", userID).
			Update("is_primary", false).Error; err != nil {
			return err
		}

		var nextMethod entity.PayoutMethod
		err := tx.Where("user_id = ?", userID).Order("created_at ASC").Order("id ASC").First(&nextMethod).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		if err != nil {
			return err
		}
		return tx.Model(&nextMethod).Update("is_primary", true).Error
	})
}

func (r *PayoutMethodRepository) SetPrimary(ctx context.Context, userID, methodID uuid.UUID) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		var user entity.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&user, "id = ?", userID).Error; err != nil {
			return err
		}
		var method entity.PayoutMethod
		if err := tx.Where("user_id = ? AND id = ?", userID, methodID).First(&method).Error; err != nil {
			return err
		}
		if err := tx.Model(&entity.PayoutMethod{}).
			Where("user_id = ?", userID).
			Update("is_primary", false).Error; err != nil {
			return err
		}
		return tx.Model(&method).Update("is_primary", true).Error
	})
}
