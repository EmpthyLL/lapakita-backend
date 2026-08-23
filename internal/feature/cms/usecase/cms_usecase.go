package usecase

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/cms/dto"
	"lapakita-backend/internal/feature/cms/repository"
)

type CMSUsecase struct {
	cmsRepo *repository.CMSRepository
}

func NewCMSUsecase(cmsRepo *repository.CMSRepository) *CMSUsecase {
	return &CMSUsecase{
		cmsRepo: cmsRepo,
	}
}

func (u *CMSUsecase) GetGroupedFAQs(ctx context.Context, lang string, roleType string, req *dto.FAQQueryRequest) ([]dto.FAQCategoryResponse, error) {
	activeLang := req.Lang
	if activeLang == "" {
		activeLang = lang
	}

	if roleType == "" {
		roleType = "all"
	}

	faqs, err := u.cmsRepo.GetFAQs(ctx, activeLang, roleType)
	if err != nil {
		return nil, err
	}

	catMap := make(map[string]*dto.FAQCategoryResponse)
	var catOrder []string

	for _, faq := range faqs {
		catID := faq.RoleType
		if _, exists := catMap[catID]; !exists {
			catMap[catID] = &dto.FAQCategoryResponse{
				ID:        catID,
				SubTopics: []dto.FAQSubTopic{},
			}
			catOrder = append(catOrder, catID)
		}

		cat := catMap[catID]

		if cat.LastUpdatedAt == nil || faq.UpdatedAt.After(*cat.LastUpdatedAt) {
			t := faq.UpdatedAt
			cat.LastUpdatedAt = &t
		}

		subTopicIdx := -1
		for i, st := range cat.SubTopics {
			if st.Title == faq.SubTopicTitle {
				subTopicIdx = i
				break
			}
		}

		item := dto.FAQItem{
			ID:       faq.ID.String(),
			Question: faq.Question,
			Answer:   faq.Answer,
		}

		if subTopicIdx == -1 {
			cat.SubTopics = append(cat.SubTopics, dto.FAQSubTopic{
				Title: faq.SubTopicTitle,
				Items: []dto.FAQItem{item},
			})
		} else {
			cat.SubTopics[subTopicIdx].Items = append(cat.SubTopics[subTopicIdx].Items, item)
		}
	}

	var result []dto.FAQCategoryResponse
	for _, catID := range catOrder {
		result = append(result, *catMap[catID])
	}

	return result, nil
}

func (u *CMSUsecase) GetLegalDocument(ctx context.Context, lang string, docType string, req *dto.LegalDocumentQueryRequest) (*dto.LegalDocumentResponse, error) {
	activeLang := req.Lang
	if activeLang == "" {
		activeLang = lang
	}

	doc, err := u.cmsRepo.GetLegalDocument(ctx, docType, activeLang)
	if err != nil || doc == nil {
		return nil, err
	}

	var sections []dto.LegalSection
	if err := json.Unmarshal(doc.SectionsJSON, &sections); err != nil {
		return nil, err
	}

	return &dto.LegalDocumentResponse{
		LastUpdatedAt: &doc.UpdatedAt,
		Data:          sections,
	}, nil
}
