package usecase

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/repository"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/jwt"
	"lapakita-backend/pkg/mailer"
	"lapakita-backend/pkg/storage"

	"github.com/redis/go-redis/v9"
)

type AuthUsecase struct {
	Repo       *repository.AuthRepository
	RDB        *redis.Client
	JWTService *jwt.JWTService
	Mailer     *mailer.Mailer
	Seaweed    *storage.SeaweedFSService
}

func NewAuthUsecase(repo *repository.AuthRepository, rdb *redis.Client, jwtService *jwt.JWTService, mailer *mailer.Mailer, seaweed *storage.SeaweedFSService) *AuthUsecase {
	return &AuthUsecase{Repo: repo, RDB: rdb, JWTService: jwtService, Mailer: mailer, Seaweed: seaweed}
}

func GenerateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

func BuildUserPayload(user *entity.User) dto.UserPayload {
	defaultPhone := dto.PhoneNumber{DialCode: user.PhoneNumbers.GetPrimaryDialCode(), Number: user.PhoneNumbers.GetPrimaryNumber()}
	var defaultAvatarPtr *string
	if user.DefaultAvatarURL != nil && *user.DefaultAvatarURL != "" {
		defaultAvatarPtr = user.DefaultAvatarURL
	}
	activePlan := user.SubscriptionPlan
	var subExpiresAtPtr *string
	if user.SubscriptionExpiresAt != nil {
		if time.Now().After(*user.SubscriptionExpiresAt) {
			activePlan = "free"
		} else {
			formatted := user.SubscriptionExpiresAt.Format(time.RFC3339)
			subExpiresAtPtr = &formatted
		}
	}
	personas := make(map[string]dto.PersonaDetail)
	for _, phone := range user.PhoneNumbers {
		for _, role := range phone.Roles {
			cleanRole := strings.TrimSpace(role)
			if cleanRole == "" {
				continue
			}
			displayName, avatarURL := "", ""
			if profile, exists := user.RoleProfiles[cleanRole]; exists {
				displayName, avatarURL = profile.DisplayName, profile.AvatarURL
			}
			personas[cleanRole] = dto.PersonaDetail{DisplayName: displayName, AvatarURL: avatarURL, Phone: dto.PhoneNumber{DialCode: phone.DialCode, Number: phone.Number}}
		}
	}
	return dto.UserPayload{
		ID: api.EncodeResponse(user.ID), DefaultName: user.Name, DefaultAvatarURL: defaultAvatarPtr,
		DefaultPhone: defaultPhone, Email: user.Email, IsPasswordSet: user.PasswordHash != "",
		ActiveRole: user.ActiveRole, SubscriptionPlan: activePlan, SubscriptionExpiresAt: subExpiresAtPtr, Personas: personas,
	}
}

func BuildAuthResponse(jwtService *jwt.JWTService, user *entity.User) (*dto.AuthResponseData, error) {
	userPayload := BuildUserPayload(user)
	tokenPair, err := jwtService.GenerateTokenPair(jwt.JWTCustomClaims{UserPayload: userPayload})
	if err != nil {
		return nil, err
	}
	userPayload.Token = tokenPair.AccessToken
	return &dto.AuthResponseData{User: userPayload, AccessToken: tokenPair.AccessToken, RefreshToken: tokenPair.RefreshToken}, nil
}
