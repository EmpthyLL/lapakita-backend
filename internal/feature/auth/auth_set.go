package auth

import (
	"lapakita-backend/internal/feature/auth/handler"
	"lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/internal/feature/auth/usecase"

	"github.com/google/wire"
)

var AuthFeatureSet = wire.NewSet(
	repository.NewAuthRepository,
	usecase.NewAuthUsecase,
	handler.NewAuthHandler,
)
