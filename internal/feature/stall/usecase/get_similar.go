package usecase

import (
	"context"
	"errors"
	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/pkg/api"
)

func (u *StallUsecase) GetSimilar(ctx context.Context, req dto.GetSimilarStallsRequest) ([]dto.CompactStallResponse, api.PaginationMeta, error) {
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
		responseList = append(responseList, MapToCompactResponse(&stall))
	}
	return responseList, meta, nil
}
