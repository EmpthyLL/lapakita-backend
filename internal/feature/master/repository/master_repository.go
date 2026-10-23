package repository

import (
	"context"
	"fmt"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

	"gorm.io/gorm"
)

type MasterRepository struct {
	db *gorm.DB
}

func NewMasterRepository(db *gorm.DB) *MasterRepository {
	return &MasterRepository{db: db}
}

func (r *MasterRepository) GetFAQs(ctx context.Context, lang, roleType string) ([]entity.CMSPublicFAQ, error) {
	var faqs []entity.CMSPublicFAQ
	query := r.db.WithContext(ctx).Model(&entity.CMSPublicFAQ{})

	if lang != "" {
		query = query.Where("lang = ?", lang)
	}
	if roleType != "" {
		query = query.Where("role_type = ?", roleType)
	}

	if err := query.Order("sort_order ASC, created_at ASC").Find(&faqs).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}

func (r *MasterRepository) GetLegalDocument(ctx context.Context, docType, lang string) (*entity.CMSLegalDocument, error) {
	var doc entity.CMSLegalDocument
	err := r.db.WithContext(ctx).
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

func (r *MasterRepository) CreateContactInquiry(ctx context.Context, inquiry *entity.ContactInquiry) error {
	return database.RunInTransaction(ctx, r.db, func(tx *gorm.DB) error {
		return tx.Create(inquiry).Error
	})
}

func (r *MasterRepository) GetBusinessTypes(ctx context.Context, lang string, req *dto.GetBusinessTypesRequest) ([]entity.BusinessType, api.PaginationMeta, error) {
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

	orderClause := fmt.Sprintf("label_lang->>'%s' ASC", lang)
	if err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"business_types",
		&businessTypes,
		&meta,
		orderClause,
	); err != nil {
		return nil, api.PaginationMeta{}, err
	}
	return businessTypes, meta, nil
}
