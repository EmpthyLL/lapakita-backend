package public

import (
	"lapakita-backend/internal/feature/public/handler"
	"lapakita-backend/internal/feature/public/repository"
	"lapakita-backend/internal/feature/public/usecase"

	"github.com/google/wire"
)

var PublicFeatureSet = wire.NewSet(
	repository.NewPublicRepository,
	usecase.NewPublicUsecase,
	handler.NewPublicHandler,
)
