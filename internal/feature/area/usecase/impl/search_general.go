package usecase

import (
	"context"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"
)

func (u *AreaUsecase) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchGeneral(ctx, req)
}
