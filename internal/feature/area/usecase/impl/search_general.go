package impl

import (
	"context"

	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/internal/feature/area/usecase"
	"lapakita-backend/pkg/api"
)

func SearchGeneral(ctx context.Context, u *usecase.AreaUsecase, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.GeoClient.SearchGeneral(ctx, req)
}
