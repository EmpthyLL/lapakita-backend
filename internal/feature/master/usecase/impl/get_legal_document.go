package impl

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/internal/feature/master/usecase"
)

func GetLegalDocument(ctx context.Context, u *usecase.MasterUsecase, lang string, docType string, req *dto.LegalDocumentQueryRequest) (*dto.LegalDocumentResponse, error) {
	activeLang := usecase.ResolveLanguage(req.Lang, lang)

	doc, err := u.Repo.GetLegalDocument(ctx, docType, activeLang)
	if err != nil || doc == nil {
		return nil, err
	}

	var sections []dto.LegalSection
	if err := json.Unmarshal(doc.SectionsJSON, &sections); err != nil {
		return nil, err
	}

	return &dto.LegalDocumentResponse{LastUpdatedAt: &doc.UpdatedAt, Data: sections}, nil
}
