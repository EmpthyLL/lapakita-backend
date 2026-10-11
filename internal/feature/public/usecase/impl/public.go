package usecase

import (
	"lapakita-backend/internal/feature/public/repository"
	"lapakita-backend/pkg/mailer"
)

type PublicUsecase struct {
	cmsRepo *repository.PublicRepository
	mailer  *mailer.Mailer
}

func NewPublicUsecase(cmsRepo *repository.PublicRepository, mailer *mailer.Mailer) *PublicUsecase {
	return &PublicUsecase{cmsRepo: cmsRepo, mailer: mailer}
}
