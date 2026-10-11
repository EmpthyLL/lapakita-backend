package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
	"google.golang.org/api/idtoken"
)

func (u *AuthUsecase) GoogleAuth(ctx context.Context, req dto.GoogleAuthRequest) (*dto.AuthResponseData, error) {
	payload, err := idtoken.Validate(ctx, req.IDToken, "")
	if err != nil {
		return nil, errors.New(string(i18n.KeyGoogleAuthFailed))
	}
	email := fmt.Sprintf("%v", payload.Claims["email"])
	name := fmt.Sprintf("%v", payload.Claims["name"])
	picture := fmt.Sprintf("%v", payload.Claims["picture"])
	user, _ := u.repo.FindUserByEmail(ctx, email)
	uploadAvatar := func(rawSource string) string {
		cleanSource := strings.TrimSpace(rawSource)
		if cleanSource == "" {
			return ""
		}
		fileName := fmt.Sprintf("avatar_%s.jpg", uuid.New().String())
		seaweedURL, err := u.seaweed.UploadFromURL(ctx, cleanSource, fileName, "/avatars")
		if err != nil {
			fmt.Printf("[SeaweedFS Upload Error]: %v\n", err)
			return cleanSource
		}
		return seaweedURL
	}
	if user == nil {
		var avatarPtr *string
		if picture != "" {
			if seaweedURL := uploadAvatar(picture); seaweedURL != "" {
				avatarPtr = &seaweedURL
			}
		}
		newUser := &entity.User{Name: name, Email: email, PasswordHash: "", DefaultAvatarURL: avatarPtr, ActiveRole: "tenant", SubscriptionPlan: "free", PhoneNumbers: entity.PhoneNumbers{}, RoleProfiles: entity.RoleProfiles{}}
		if err := u.repo.CreateUser(ctx, newUser); err != nil {
			return nil, err
		}
		user = newUser
	} else {
		needsUpload := user.DefaultAvatarURL == nil || *user.DefaultAvatarURL == "" || strings.Contains(*user.DefaultAvatarURL, "googleusercontent.com") || strings.HasPrefix(*user.DefaultAvatarURL, "data:image") || (!strings.HasPrefix(*user.DefaultAvatarURL, "http://") && !strings.HasPrefix(*user.DefaultAvatarURL, "https://"))
		if needsUpload && picture != "" {
			seaweedURL := uploadAvatar(picture)
			if seaweedURL != "" && (user.DefaultAvatarURL == nil || seaweedURL != *user.DefaultAvatarURL) {
				user.DefaultAvatarURL = &seaweedURL
				_ = u.repo.UpdateUser(ctx, user)
			}
		}
	}
	return usecase.BuildAuthResponse(u.jwtService, user)
}
