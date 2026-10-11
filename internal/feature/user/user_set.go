package user

import (
	"lapakita-backend/internal/feature/user/handler"
	"lapakita-backend/internal/feature/user/repository"
	usecase "lapakita-backend/internal/feature/user/usecase/impl"

	"github.com/google/wire"
)

var UserFeatureSet = wire.NewSet(
	repository.NewUserRepository,
	usecase.NewUserUsecase,
	handler.NewUserHandler,
)
