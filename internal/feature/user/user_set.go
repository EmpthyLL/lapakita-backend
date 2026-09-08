package user

import (
	"lapakita-backend/internal/feature/user/handler"
	"lapakita-backend/internal/feature/user/repository"
	"lapakita-backend/internal/feature/user/usecase"

	"github.com/google/wire"
)

var UserFeatureSet = wire.NewSet(
	repository.NewUserRepository,
	usecase.NewUserUsecase,
	handler.NewUserHandler,
)
