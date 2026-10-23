package usecase

import "lapakita-backend/internal/feature/wallet/repository"

type WalletUsecase struct {
	Repo *repository.WalletRepository
}

func NewWalletUsecase(repo *repository.WalletRepository) *WalletUsecase {
	return &WalletUsecase{Repo: repo}
}
