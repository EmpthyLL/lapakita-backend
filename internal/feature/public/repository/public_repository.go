package repository

import (
	"context"

	"lapakita-backend/internal/entity"

	"gorm.io/gorm"
)

type PublicRepository struct {
	database *gorm.DB
}

func NewPublicRepository(db *gorm.DB) *PublicRepository {
	return &PublicRepository{database: db}
}

func (r *PublicRepository) GetFAQs(ctx context.Context, lang, roleType string) ([]entity.CMSPublicFAQ, error) {
	var faqs []entity.CMSPublicFAQ
	query := r.database.WithContext(ctx).Model(&entity.CMSPublicFAQ{})

	if lang != "" {
		query = query.Where("lang = ?", lang)
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

func (r *PublicRepository) GetLegalDocument(ctx context.Context, docType, lang string) (*entity.CMSLegalDocument, error) {
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

func (r *PublicRepository) CreateContactInquiry(ctx context.Context, inquiry *entity.ContactInquiry) error {
	return r.database.WithContext(ctx).Create(inquiry).Error
}
