package impl

import (
	"context"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/internal/feature/master/usecase"
	"lapakita-backend/pkg/api"
)

func SearchDetail(ctx context.Context, u *usecase.MasterUsecase, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.GeoClient.SearchDetail(ctx, req)
}
