package repository

import (
	"context"

	"lapakita-backend/internal/entity"

	"gorm.io/gorm"
)

type CMSRepository struct {
	database *gorm.DB
}

func NewCMSRepository(db *gorm.DB) *CMSRepository {
	return &CMSRepository{database: db}
}

func (r *CMSRepository) GetFAQs(ctx context.Context, lang, category, roleType string) ([]entity.CMSPublicFAQ, error) {
	var faqs []entity.CMSPublicFAQ
	query := r.database.WithContext(ctx).Model(&entity.CMSPublicFAQ{})

	if lang != "" {
		query = query.Where("lang = ?", lang)
	}
	if category != "" {
		query = query.Where("category_id = ?", category)
	}
	if roleType != "" {
		query = query.Where("role_type = ?", roleType)
	}

	err := query.Order("sort_order ASC, created_at ASC").Find(&faqs).Error
	if err != nil {
		return nil, err
	}
	return faqs, nil
}

func (r *CMSRepository) GetLegalDocument(ctx context.Context, docType, lang string) (*entity.CMSLegalDocument, error) {
	var doc entity.CMSLegalDocument
	err := r.database.WithContext(ctx).
		Where("doc_type = ? AND lang = ?", docType, lang).
		First(&doc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &doc, nil
}
