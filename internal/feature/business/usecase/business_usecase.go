package usecase

import (
	"lapakita-backend/internal/feature/business/repository"
	"lapakita-backend/pkg/storage"
)

type BusinessUsecase struct {
	businessRepo *repository.BusinessRepository
	seaweed      *storage.SeaweedFSService
}

func NewBusinessUsecase(businessRepo *repository.BusinessRepository, seaweed *storage.SeaweedFSService) *BusinessUsecase {
	return &BusinessUsecase{
		businessRepo: businessRepo,
		seaweed:      seaweed,
	}
}
