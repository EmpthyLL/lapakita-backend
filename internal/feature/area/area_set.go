package area

import (
	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/handler"
	"lapakita-backend/internal/feature/area/usecase"

	"github.com/google/wire"
)

var AreaFeatureSet = wire.NewSet(
	client.NewGeoapifyClient,
	usecase.NewAreaUsecase,
	handler.NewAreaHandler,
)
