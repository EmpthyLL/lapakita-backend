package wallet

import (
	"lapakita-backend/internal/feature/wallet/handler"
	"lapakita-backend/internal/feature/wallet/repository"
	"lapakita-backend/internal/feature/wallet/usecase"

	"github.com/google/wire"
)

var PayoutMethodFeatureSet = wire.NewSet(
	repository.NewPayoutMethodRepository,
	usecase.NewPayoutMethodUsecase,
	handler.NewPayoutMethodHandler,
)
