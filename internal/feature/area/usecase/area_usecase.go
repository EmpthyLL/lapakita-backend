package usecase

import (
	"context"
	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/dto"
)

type AreaUsecase interface {
	GetNearbyPlaces(ctx context.Context, req dto.GetAreaRequest) ([]dto.AreaResponse, error)
}

type areaUsecase struct {
	geoClient client.GeoapifyClient
}

func NewAreaUsecase(geoClient client.GeoapifyClient) AreaUsecase {
	return &areaUsecase{
		geoClient: geoClient,
	}
}

func (u *areaUsecase) GetNearbyPlaces(ctx context.Context, req dto.GetAreaRequest) ([]dto.AreaResponse, error) {
	// Di layer ini kamu bisa menyelipkan business logic opsional jika ada
	return u.geoClient.FetchPlaces(ctx, req)
}
