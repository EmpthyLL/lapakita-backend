package wallet

import (
	"lapakita-backend/internal/feature/wallet/handler"
	"lapakita-backend/internal/feature/wallet/repository"
	"lapakita-backend/internal/feature/wallet/usecase"

	"github.com/google/wire"
)

var WalletFeatureSet = wire.NewSet(
	repository.NewWalletRepository,
	usecase.NewWalletUsecase,
	handler.NewWalletHandler,
)
