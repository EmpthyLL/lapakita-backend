package usecase

import (
	"lapakita-backend/internal/feature/area/client"

	"github.com/redis/go-redis/v9"
)

type AreaUsecase struct {
	geoClient *client.GeoapifyClient
	rdb       *redis.Client
}

func NewAreaUsecase(geoClient *client.GeoapifyClient, rdb *redis.Client) *AreaUsecase {
	return &AreaUsecase{geoClient: geoClient, rdb: rdb}
}
