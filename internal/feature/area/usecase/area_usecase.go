package usecase

import (
	"context"

	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"
)

type AreaUsecase interface {
	SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error)
	SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error)
}

type areaUsecase struct {
	geoClient client.GeoapifyClient
}

func NewAreaUsecase(geoClient client.GeoapifyClient) AreaUsecase {
	return &areaUsecase{
		geoClient: geoClient,
	}
}

func (u *areaUsecase) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchGeneral(ctx, req)
}

func (u *areaUsecase) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchDetail(ctx, req)
}
