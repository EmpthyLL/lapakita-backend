package repository

import (
	"context"
	"errors"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/wallet/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) Exists(ctx context.Context, userID uuid.UUID, payoutType, providerCode, accountNumber string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&entity.PayoutMethod{}).
		Where("user_id = ? AND LOWER(payout_type) = LOWER(?) AND LOWER(provider_code) = LOWER(?) AND TRIM(account_number) = ?",
			userID, payoutType, providerCode, strings.TrimSpace(accountNumber)).
		Count(&count).Error
	return count > 0, err
}

func (r *WalletRepository) GetByUserID(ctx context.Context, userID uuid.UUID, req dto.GetPayoutMethodsRequest) ([]entity.PayoutMethod, api.PaginationMeta, error) {
	var methods []entity.PayoutMethod
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).
		Model(&entity.PayoutMethod{}).
		Where("user_id = ?", userID)

	// Filter Search Gabungan: Nomor Rekening / Nama Pemilik Rekening
	if req.Search != "" {
		searchPattern := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("(LOWER(account_number) LIKE ? OR LOWER(account_holder_name) LIKE ?)", searchPattern, searchPattern)
	}

	// Filter Provider Gabungan: Code (e.g. "BCA") / Name (e.g. "Bank Central Asia")
	if req.Provider != "" {
		providerPattern := "%" + strings.ToLower(req.Provider) + "%"
		query = query.Where("(LOWER(provider_code) LIKE ? OR LOWER(provider_name) LIKE ?)", providerPattern, providerPattern)
	}

	// Filter Payout Type (bank_account / e_wallet)
	if req.PayoutType != "" {
		query = query.Where("payout_type = ?", req.PayoutType)
	}

	// Filter Is Primary
	if req.IsPrimary != nil {
		query = query.Where("is_primary = ?", *req.IsPrimary)
	}

	// Order: utamakan metode primary terlebih dahulu, kemudian tanggal dibuat paling baru
	orderClause := "is_primary DESC, created_at DESC"

	if err := database.AutoPaginate(query, req.BasePaginationRequest, "payout_methods", &methods, &meta, orderClause); err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return methods, meta, nil
}

func (r *WalletRepository) Create(ctx context.Context, method *entity.PayoutMethod) error {
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

func (r *WalletRepository) Delete(ctx context.Context, userID, methodID uuid.UUID) error {
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

func (r *WalletRepository) SetPrimary(ctx context.Context, userID, methodID uuid.UUID) error {
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
