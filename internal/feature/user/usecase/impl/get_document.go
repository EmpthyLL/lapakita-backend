package impl

import (
	"context"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/api"

	"github.com/google/uuid"
)

func GetDocument(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, req dto.GetDocumentRequest) ([]dto.GetDocumentResponse, api.PaginationMeta, error) {
	docs, meta, err := u.Repo.GetDocument(ctx, userID, &req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}
	res := make([]dto.GetDocumentResponse, 0, len(docs))
	for _, doc := range docs {
		res = append(res, dto.GetDocumentResponse{ID: api.EncodeResponse(doc.ID), CountryCode: doc.CountryCode, DocumentType: doc.DocumentType, DocumentLabel: doc.DocumentLabel, FullNameIdentity: doc.FullNameIdentity, DocumentNumber: doc.DocumentNumber, DocumentPhotoURL: doc.DocumentPhotoURL})
	}
	return res, meta, nil
}
