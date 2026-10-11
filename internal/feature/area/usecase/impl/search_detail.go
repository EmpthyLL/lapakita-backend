package usecase

import (
	"context"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"
)

func (u *AreaUsecase) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchDetail(ctx, req)
}
