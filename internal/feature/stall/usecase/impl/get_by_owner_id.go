package impl

import (
	"context"
	"lapakita-backend/internal/feature/stall/dto"
	usecase "lapakita-backend/internal/feature/stall/usecase"
	"lapakita-backend/pkg/api"
)

func GetByOwnerID(ctx context.Context, u *usecase.StallUsecase, req dto.GetOwnerStallsRequest) ([]dto.CompactStallResponse, api.PaginationMeta, error) {
	stalls, meta, err := u.StallRepository.FindByOwnerID(ctx, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}
	responseList := make([]dto.CompactStallResponse, 0, len(stalls))
	for _, stall := range stalls {
		responseList = append(responseList, usecase.MapToCompactResponse(&stall))
	}
	return responseList, meta, nil
}
