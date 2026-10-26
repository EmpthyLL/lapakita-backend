package usecase

import (
	"context"
	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/pkg/api"
)

func (u *StallUsecase) Search(ctx context.Context, req dto.SearchStallRequest) ([]dto.CompactStallResponse, api.PaginationMeta, error) {
	stalls, meta, err := u.StallRepository.Search(ctx, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}
	responseList := make([]dto.CompactStallResponse, 0, len(stalls))
	for _, stall := range stalls {
		responseList = append(responseList, MapToCompactResponse(&stall))
	}
	return responseList, meta, nil
}
