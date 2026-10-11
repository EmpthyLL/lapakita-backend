package stall

import (
	"lapakita-backend/internal/feature/stall/handler"
	"lapakita-backend/internal/feature/stall/repository"
	usecase "lapakita-backend/internal/feature/stall/usecase/impl"

	"github.com/google/wire"
)

var StallFeatureSet = wire.NewSet(
	repository.NewStallRepository,
	usecase.NewStallUsecase,
	handler.NewStallHandler,
)
