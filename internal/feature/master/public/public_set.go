package public

import (
	"lapakita-backend/internal/feature/master/public/handler"
	"lapakita-backend/internal/feature/master/public/repository"
	"lapakita-backend/internal/feature/master/public/usecase"

	"github.com/google/wire"
)

var PublicFeatureSet = wire.NewSet(
	repository.NewPublicRepository,
	usecase.NewPublicUsecase,
	handler.NewPublicHandler,
)
