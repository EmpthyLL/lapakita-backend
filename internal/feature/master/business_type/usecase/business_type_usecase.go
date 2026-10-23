package usecase

import "lapakita-backend/internal/feature/master/business_type/repository"

type BusinessTypeUsecase struct {
	Repo *repository.BusinessTypeRepository
}

func NewBusinessTypeUsecase(repo *repository.BusinessTypeRepository) *BusinessTypeUsecase {
	return &BusinessTypeUsecase{Repo: repo}
}
