package cms

import (
	"lapakita-backend/internal/feature/cms/handler"
	"lapakita-backend/internal/feature/cms/repository"
	"lapakita-backend/internal/feature/cms/usecase"

	"github.com/google/wire"
)

var CMSFeatureSet = wire.NewSet(
	repository.NewCMSRepository,
	usecase.NewCMSUsecase,
	handler.NewCMSHandler,
)
