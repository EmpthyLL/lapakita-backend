package usecase

import "lapakita-backend/internal/feature/wallet/repository"

type PayoutMethodUsecase struct {
	Repo *repository.PayoutMethodRepository
}

func NewPayoutMethodUsecase(repo *repository.PayoutMethodRepository) *PayoutMethodUsecase {
	return &PayoutMethodUsecase{Repo: repo}
}
