package usecase

import (
	"context"

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

func (u *CMSUsecase) GetGroupedFAQs(ctx context.Context, lang string, req *dto.FAQQueryReq) ([]dto.FAQCategoryResponse, error) {
	// Gunakan lang dari request query jika diisi, jika tidak gunakan lang dari header
	activeLang := req.Lang
	if activeLang == "" {
		activeLang = lang
	}

	faqs, err := u.cmsRepo.GetFAQs(ctx, activeLang, req.Category, req.RoleType)
	if err != nil {
		return nil, err
	}

	catMap := make(map[string]*dto.FAQCategoryResponse)
	var catOrder []string

	for _, faq := range faqs {
		catID := faq.CategoryID
		if _, exists := catMap[catID]; !exists {
			catMap[catID] = &dto.FAQCategoryResponse{
				ID:          catID,
				Label:       getCategoryLabel(catID, activeLang),
				Description: getCategoryDescription(catID, activeLang),
				SubTopics:   []dto.FAQSubTopic{},
			}
			catOrder = append(catOrder, catID)
		}

		cat := catMap[catID]
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

func (u *CMSUsecase) GetLegalDocument(ctx context.Context, lang string, req *dto.LegalDocumentQueryReq) (*dto.LegalDocumentResponse, error) {
	activeLang := req.Lang
	if activeLang == "" {
		activeLang = lang
	}

	doc, err := u.cmsRepo.GetLegalDocument(ctx, req.DocType, activeLang)
	if err != nil || doc == nil {
		return nil, err
	}

	return &dto.LegalDocumentResponse{
		DocType:  doc.DocType,
		Label:    getLegalLabel(doc.DocType),
		Title:    doc.Title,
		Intro:    doc.Intro,
		Sections: doc.SectionsJSON,
	}, nil
}

func getCategoryLabel(id string, lang string) string {
	switch id {
	case "all":
		if lang == "id" {
			return "Umum & Platform"
		}
		return "General & Platform"
	case "tenant":
		if lang == "id" {
			return "Penyewa & Bisnis"
		}
		return "Tenant & Business"
	case "owner":
		if lang == "id" {
			return "Pemilik Lapak"
		}
		return "Stall Owner"
	case "supplier":
		return "Supplier & B2B"
	default:
		return id
	}
}

func getCategoryDescription(id string, lang string) string {
	switch id {
	case "all":
		if lang == "id" {
			return "Pelajari tentang ekosistem Lapakita, akun tunggal multi-peran, dan paket harga."
		}
		return "Learn about Lapakita's ecosystem, single account multi-role, and pricing plans."
	case "tenant":
		if lang == "id" {
			return "Panduan untuk pelaku usaha yang menjalankan POS, mencari lapak, dan menganalisis keuangan."
		}
		return "Guides for business operators running POS, finding stalls, and analyzing finances."
	case "owner":
		if lang == "id" {
			return "Informasi untuk pemilik properti dalam mengelola listing, penyewa, dan deposit escrow."
		}
		return "Information for property owners managing listings, tenants, and escrow deposits."
	case "supplier":
		if lang == "id" {
			return "Detail untuk grosir dan distributor yang terhubung dengan pembeli UMKM."
		}
		return "Details for wholesalers and distributors connecting with SME buyers."
	default:
		return ""
	}
}

func getLegalLabel(docType string) string {
	switch docType {
	case "terms":
		return "Terms & Conditions"
	case "privacy":
		return "Privacy Policy"
	case "cookies":
		return "Cookies Policy"
	default:
		return docType
	}
}
