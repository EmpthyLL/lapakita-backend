package businesstype

import (
	"lapakita-backend/internal/feature/business_type/handler"
	"lapakita-backend/internal/feature/business_type/repository"
	usecase "lapakita-backend/internal/feature/business_type/usecase/impl"

	"github.com/google/wire"
)

var BusinessTypeFeatureSet = wire.NewSet(
	repository.NewBusinessTypeRepository,
	usecase.NewBusinessTypeUsecase,
	handler.NewBusinessTypeHandler,
)
