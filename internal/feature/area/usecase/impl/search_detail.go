package impl

import (
	"context"

	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/internal/feature/area/usecase"
	"lapakita-backend/pkg/api"
)

func SearchDetail(ctx context.Context, u *usecase.AreaUsecase, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.GeoClient.SearchDetail(ctx, req)
}
