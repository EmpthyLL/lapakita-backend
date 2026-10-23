package impl

import (
	"context"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/internal/feature/master/usecase"
	"lapakita-backend/pkg/api"
)

func SearchGeneral(ctx context.Context, u *usecase.MasterUsecase, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.GeoClient.SearchGeneral(ctx, req)
}
