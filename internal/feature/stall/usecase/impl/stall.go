package usecase

import (
	authRepo "lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/internal/feature/stall/repository"
	"lapakita-backend/pkg/storage"
)

type StallUsecase struct {
	repo     *repository.StallRepository
	userRepo *authRepo.AuthRepository
	seaweed  *storage.SeaweedFSService
}

func NewStallUsecase(repo *repository.StallRepository, userRepo *authRepo.AuthRepository, seaweed *storage.SeaweedFSService) *StallUsecase {
	return &StallUsecase{repo: repo, userRepo: userRepo, seaweed: seaweed}
}
