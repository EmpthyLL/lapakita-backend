package usecase

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/public/dto"
	sharedusecase "lapakita-backend/internal/feature/public/usecase"
)

func (u *PublicUsecase) GetLegalDocument(ctx context.Context, lang string, docType string, req *dto.LegalDocumentQueryRequest) (*dto.LegalDocumentResponse, error) {
	activeLang := sharedusecase.ResolveLanguage(req.Lang, lang)

	doc, err := u.cmsRepo.GetLegalDocument(ctx, docType, activeLang)
	if err != nil || doc == nil {
		return nil, err
	}

	var sections []dto.LegalSection
	if err := json.Unmarshal(doc.SectionsJSON, &sections); err != nil {
		return nil, err
	}

	return &dto.LegalDocumentResponse{LastUpdatedAt: &doc.UpdatedAt, Data: sections}, nil
}
