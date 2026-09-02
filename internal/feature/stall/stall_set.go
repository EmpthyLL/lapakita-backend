package stall

import (
	"lapakita-backend/internal/feature/stall/handler"
	"lapakita-backend/internal/feature/stall/repository"
	"lapakita-backend/internal/feature/stall/usecase"

	"github.com/google/wire"
)

var StallFeatureSet = wire.NewSet(
	repository.NewStallRepository,
	usecase.NewStallUsecase,
	handler.NewStallHandler,
)
