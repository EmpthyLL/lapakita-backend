package impl

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/master/public/dto"
	"lapakita-backend/internal/feature/master/public/usecase"
	sharedusecase "lapakita-backend/internal/feature/master/public/usecase"
)

func GetLegalDocument(ctx context.Context, u *usecase.PublicUsecase, lang string, docType string, req *dto.LegalDocumentQueryRequest) (*dto.LegalDocumentResponse, error) {
	activeLang := sharedusecase.ResolveLanguage(req.Lang, lang)

	doc, err := u.CMSRepo.GetLegalDocument(ctx, docType, activeLang)
	if err != nil || doc == nil {
		return nil, err
	}

	var sections []dto.LegalSection
	if err := json.Unmarshal(doc.SectionsJSON, &sections); err != nil {
		return nil, err
	}

	return &dto.LegalDocumentResponse{LastUpdatedAt: &doc.UpdatedAt, Data: sections}, nil
}
