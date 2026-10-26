package usecase

import (
	"context"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
)

func (u *MasterUsecase) GetGroupedFAQs(ctx context.Context, lang string, roleType string, req *dto.FAQQueryRequest) ([]dto.FAQCategoryResponse, error) {
	activeLang := ResolveLanguage(req.Lang, lang)
	if roleType == "" {
		roleType = "all"
	}

	faqs, err := u.Repo.GetFAQs(ctx, activeLang, roleType)
	if err != nil {
		return nil, err
	}

	catMap := make(map[string]*dto.FAQCategoryResponse)
	var catOrder []string

	for _, faq := range faqs {
		catID := faq.RoleType
		if _, exists := catMap[catID]; !exists {
			catMap[catID] = &dto.FAQCategoryResponse{ID: catID, SubTopics: []dto.FAQSubTopic{}}
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

		item := dto.FAQItem{ID: api.EncodeResponse(faq.ID), Question: faq.Question, Answer: faq.Answer}
		if subTopicIdx == -1 {
			cat.SubTopics = append(cat.SubTopics, dto.FAQSubTopic{Title: faq.SubTopicTitle, Items: []dto.FAQItem{item}})
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
