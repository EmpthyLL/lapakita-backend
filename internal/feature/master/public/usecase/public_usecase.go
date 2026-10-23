package usecase

import (
	"lapakita-backend/internal/feature/master/public/repository"
	"lapakita-backend/pkg/mailer"
)

type PublicUsecase struct {
	CMSRepo *repository.PublicRepository
	Mailer  *mailer.Mailer
}

func NewPublicUsecase(cmsRepo *repository.PublicRepository, mailer *mailer.Mailer) *PublicUsecase {
	return &PublicUsecase{CMSRepo: cmsRepo, Mailer: mailer}
}

// ResolveLanguage chooses a valid request language, falling back to the
// language supplied by the caller and then to English.
func ResolveLanguage(requested, fallback string) string {
	if requested != "" {
		return requested
	}
	if fallback != "" {
		return fallback
	}
	return "en"
}
