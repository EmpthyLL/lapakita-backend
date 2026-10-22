package impl

import (
	"context"
	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/pkg/api"

	usecase "lapakita-backend/internal/feature/stall/usecase"
)

func Search(ctx context.Context, u *usecase.StallUsecase, req dto.SearchStallRequest) ([]dto.CompactStallResponse, api.PaginationMeta, error) {
	stalls, meta, err := u.StallRepository.Search(ctx, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}
	responseList := make([]dto.CompactStallResponse, 0, len(stalls))
	for _, stall := range stalls {
		responseList = append(responseList, usecase.MapToCompactResponse(&stall))
	}
	return responseList, meta, nil
}
