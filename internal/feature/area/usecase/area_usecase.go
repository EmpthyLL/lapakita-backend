package usecase

import (
	"context"

	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/dto"
)

type AreaUsecase interface {
	SearchArea(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponse, error)
}

type areaUsecase struct {
	geoClient client.GeoapifyClient
}

func NewAreaUsecase(geoClient client.GeoapifyClient) AreaUsecase {
	return &areaUsecase{
		geoClient: geoClient,
	}
}

func (u *areaUsecase) SearchArea(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponse, error) {
	return u.geoClient.SearchAutocomplete(ctx, req, lang)
}
