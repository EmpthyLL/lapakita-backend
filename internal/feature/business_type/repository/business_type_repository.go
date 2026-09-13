package repository

import (
	"context"
	"fmt"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/business_type/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

	"gorm.io/gorm"
)

type BusinessTypeRepository struct {
	db *gorm.DB
}

func NewBusinessTypeRepository(db *gorm.DB) *BusinessTypeRepository {
	return &BusinessTypeRepository{db: db}
}

func (r *BusinessTypeRepository) GetBusinessTypes(ctx context.Context, lang string, req *dto.GetBusinessTypesRequest) ([]entity.BusinessType, api.PaginationMeta, error) {
	var businessTypes []entity.BusinessType
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).Model(&entity.BusinessType{})

	if req.Search != "" {
		searchPattern := "%" + req.Search + "%"
		query = query.Where(
			"label_lang->>'en' ILIKE ? OR label_lang->>'id' ILIKE ? OR group_name_lang->>'en' ILIKE ? OR group_name_lang->>'id' ILIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	if req.Group != "" {
		groupPattern := "%" + req.Group + "%"
		query = query.Where("group_name_lang->>'en' ILIKE ? OR group_name_lang->>'id' ILIKE ?", groupPattern, groupPattern)
	}

	// Clause pengurutan dinamis berdasarkan bahasa yang dikirim
	orderClause := fmt.Sprintf("label_lang->>'%s' ASC", lang)

	// Eksekusi AutoPaginate (otomatis memproses limit, offset, selected_id, dan metadata)
	err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"business_types",
		&businessTypes,
		&meta,
		orderClause,
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return businessTypes, meta, nil
}
