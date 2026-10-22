package impl

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
)

func CompleteProfile(ctx context.Context, u *usecase.AuthUsecase, userID uuid.UUID, req dto.CompleteProfileRequest) (*dto.UserPayload, error) {
	user, err := u.Repo.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return nil, errors.New(string(i18n.KeyUserNotFound))
	}
	user.Name = req.Name
	if req.AvatarURL != "" {
		avatarSource := strings.TrimSpace(req.AvatarURL)
		fileName := fmt.Sprintf("avatar_%s.jpg", uuid.New().String())
		seaweedURL, err := u.Seaweed.UploadFromURL(ctx, avatarSource, fileName, "/avatars")
		if err == nil && seaweedURL != "" {
			avatarSource = seaweedURL
		}
		user.DefaultAvatarURL = &avatarSource
	}
	user.PhoneNumbers = entity.PhoneNumbers{{DialCode: req.DialCode, Number: req.Phone, IsPrimary: true, Roles: []string{}, Label: "Primary Number"}}
	if err := u.Repo.UpdateUser(ctx, user); err != nil {
		return nil, err
	}
	updatedPayload := usecase.BuildUserPayload(user)
	return &updatedPayload, nil
}
