package area

import (
	"lapakita-backend/internal/feature/master/area/client"
	"lapakita-backend/internal/feature/master/area/handler"
	"lapakita-backend/internal/feature/master/area/usecase"

	"github.com/google/wire"
)

var AreaFeatureSet = wire.NewSet(
	client.NewGeoapifyClient,
	usecase.NewAreaUsecase,
	handler.NewAreaHandler,
)
