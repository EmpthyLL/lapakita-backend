package usecase

import (
	"context"

	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"
)

type AreaUsecase interface {
	SearchArea(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponseData, api.PaginationMeta, error)
}

type areaUsecase struct {
	geoClient client.GeoapifyClient
}

func NewAreaUsecase(geoClient client.GeoapifyClient) AreaUsecase {
	return &areaUsecase{
		geoClient: geoClient,
	}
}

func (u *areaUsecase) SearchArea(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchAutocomplete(ctx, req, lang)
}
