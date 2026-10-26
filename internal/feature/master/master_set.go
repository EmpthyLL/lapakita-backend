package master

import (
	"lapakita-backend/internal/feature/master/client"
	"lapakita-backend/internal/feature/master/handler"
	"lapakita-backend/internal/feature/master/repository"
	"lapakita-backend/internal/feature/master/usecase"

	"github.com/google/wire"
)

var MasterFeatureSet = wire.NewSet(
	client.NewMasterClient,
	repository.NewMasterRepository,
	usecase.NewMasterUsecase,
	handler.NewMasterHandler,
)
