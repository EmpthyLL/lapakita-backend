package usecase

import (
	"lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/mailer"
	"lapakita-backend/pkg/storage"

	"github.com/redis/go-redis/v9"
)

type AuthUsecase struct {
	repo       *repository.AuthRepository
	rdb        *redis.Client
	jwtService *jwt.JWTService
	mailer     *mailer.Mailer
	seaweed    *storage.SeaweedFSService
}

func NewAuthUsecase(repo *repository.AuthRepository, rdb *redis.Client, jwtService *jwt.JWTService, mailer *mailer.Mailer, seaweed *storage.SeaweedFSService) *AuthUsecase {
	return &AuthUsecase{repo: repo, rdb: rdb, jwtService: jwtService, mailer: mailer, seaweed: seaweed}
}
