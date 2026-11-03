package business

import (
	"lapakita-backend/internal/feature/business/handler"
	"lapakita-backend/internal/feature/business/repository"
	usecase "lapakita-backend/internal/feature/business/usecase"

	"github.com/google/wire"
)

var BusinessFeatureSet = wire.NewSet(
	repository.NewBusinessRepository,
	usecase.NewBusinessUsecase,
	handler.NewBusinessHandler,
)
