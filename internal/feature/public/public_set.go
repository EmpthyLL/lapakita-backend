package public

import (
	"lapakita-backend/internal/feature/public/handler"
	"lapakita-backend/internal/feature/public/repository"
	usecase "lapakita-backend/internal/feature/public/usecase/impl"

	"github.com/google/wire"
)

var PublicFeatureSet = wire.NewSet(
	repository.NewPublicRepository,
	usecase.NewPublicUsecase,
	handler.NewPublicHandler,
)
