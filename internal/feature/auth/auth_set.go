package auth

import (
	"lapakita-backend/internal/feature/auth/handler"
	"lapakita-backend/internal/feature/auth/repository"
	usecase "lapakita-backend/internal/feature/auth/usecase/impl"

	"github.com/google/wire"
)

var AuthFeatureSet = wire.NewSet(
	repository.NewAuthRepository,
	usecase.NewAuthUsecase,
	handler.NewAuthHandler,
)
