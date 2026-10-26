package impl

import (
	"context"
	"errors"
	"lapakita-backend/internal/feature/stall/dto"
	usecase "lapakita-backend/internal/feature/stall/usecase"
	"lapakita-backend/pkg/api"
)

func GetSimilar(ctx context.Context, u *usecase.StallUsecase, req dto.GetSimilarStallsRequest) ([]dto.CompactStallResponse, api.PaginationMeta, error) {
	targetStall, err := u.StallRepository.FindByID(ctx, req.ID)
	if err != nil || targetStall == nil {
		return nil, api.PaginationMeta{}, errors.New("stall not found")
	}
	stalls, meta, err := u.StallRepository.FindSimilar(ctx, targetStall, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}
	responseList := make([]dto.CompactStallResponse, 0, len(stalls))
	for _, stall := range stalls {
		responseList = append(responseList, usecase.MapToCompactResponse(&stall))
	}
	return responseList, meta, nil
}
