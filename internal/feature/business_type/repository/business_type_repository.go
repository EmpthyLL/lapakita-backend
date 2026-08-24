package repository

import (
	"context"
	"fmt"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/business_type/dto"

	"gorm.io/gorm"
)

type BusinessTypeRepository struct {
	db *gorm.DB
}

func NewBusinessTypeRepository(db *gorm.DB) *BusinessTypeRepository {
	return &BusinessTypeRepository{db: db}
}

func (r *BusinessTypeRepository) GetBusinessTypes(ctx context.Context, lang string, req *dto.GetBusinessTypesRequest) ([]entity.BusinessType, int64, error) {
	var businessTypes []entity.BusinessType
	var total int64

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

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.Limit
	orderClause := fmt.Sprintf("label_lang->>'%s' ASC", lang)

	err := query.Order(orderClause).
		Offset(offset).
		Limit(req.Limit).
		Find(&businessTypes).Error

	if err != nil {
		return nil, 0, err
	}

	return businessTypes, total, nil
}
