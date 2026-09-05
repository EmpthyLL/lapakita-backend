package usecase

import (
	"context"
	"encoding/json"
	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/public/dto"
	"lapakita-backend/internal/feature/public/repository"
	"lapakita-backend/pkg/mailer"
	"time"
)

type PublicUsecase struct {
	cmsRepo *repository.PublicRepository
	mailer  *mailer.Mailer
}

func NewPublicUsecase(cmsRepo *repository.PublicRepository, mailer *mailer.Mailer) *PublicUsecase {
	return &PublicUsecase{
		cmsRepo: cmsRepo,
		mailer:  mailer,
	}
}

func (u *PublicUsecase) GetGroupedFAQs(ctx context.Context, lang string, roleType string, req *dto.FAQQueryRequest) ([]dto.FAQCategoryResponse, error) {
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

func (u *PublicUsecase) GetLegalDocument(ctx context.Context, lang string, docType string, req *dto.LegalDocumentQueryRequest) (*dto.LegalDocumentResponse, error) {
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

func (u *PublicUsecase) SubmitContactInquiry(ctx context.Context, req dto.SubmitContactInquiryRequest) (*dto.ContactInquiryResponse, error) {
	inquiry := &entity.ContactInquiry{
		Name:        req.Name,
		Email:       req.Email,
		Whatsapp:    req.Whatsapp,
		Persona:     req.Persona,
		InquiryType: req.InquiryType,
		Message:     req.Message,
		Status:      "unread",
	}

	if err := u.cmsRepo.CreateContactInquiry(ctx, inquiry); err != nil {
		return nil, err
	}

	// Kirim email auto-reply secara asinkron (goroutine) agar tidak memblokir respon HTTP
	go func(email, name, persona, inqType, msg string) {
		_ = u.mailer.SendContactInquiryAutoReply(email, name, persona, inqType, msg)
	}(inquiry.Email, inquiry.Name, inquiry.Persona, inquiry.InquiryType, inquiry.Message)

	return &dto.ContactInquiryResponse{
		ID:        inquiry.ID.String(),
		Name:      inquiry.Name,
		Email:     inquiry.Email,
		Status:    inquiry.Status,
		CreatedAt: inquiry.CreatedAt.Format(time.RFC3339),
	}, nil
}
