package usecase

import (
	"lapakita-backend/internal/feature/user/repository"
	"lapakita-backend/pkg/storage"
)

type UserUsecase struct {
	repo    *repository.UserRepository
	seaweed *storage.SeaweedFSService
}

func NewUserUsecase(repo *repository.UserRepository, seaweed *storage.SeaweedFSService) *UserUsecase {
	return &UserUsecase{repo: repo, seaweed: seaweed}
}
