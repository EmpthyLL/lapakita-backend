package businesstype

import (
	"lapakita-backend/internal/feature/master/business_type/handler"
	"lapakita-backend/internal/feature/master/business_type/repository"
	"lapakita-backend/internal/feature/master/business_type/usecase"

	"github.com/google/wire"
)

var BusinessTypeFeatureSet = wire.NewSet(
	repository.NewBusinessTypeRepository,
	usecase.NewBusinessTypeUsecase,
	handler.NewBusinessTypeHandler,
)
