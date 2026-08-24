package usecase

import (
	"context"

	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"
)

type AreaUsecase struct {
	geoClient *client.GeoapifyClient
}

func NewAreaUsecase(geoClient *client.GeoapifyClient) *AreaUsecase {
	return &AreaUsecase{
		geoClient: geoClient,
	}
}

func (u *AreaUsecase) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchGeneral(ctx, req)
}

func (u *AreaUsecase) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchDetail(ctx, req)
}
