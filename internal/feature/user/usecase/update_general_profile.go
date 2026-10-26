package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func (u *UserUsecase) UpdateGeneralProfile(ctx context.Context, userID uuid.UUID, req dto.UpdateGeneralProfileRequest) (dto.GetGeneralProfileResponse, error) {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetGeneralProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}
	user.Name = req.Name
	if req.DefaultAvatarURL != nil && *req.DefaultAvatarURL != "" {
		avatarSource := strings.TrimSpace(*req.DefaultAvatarURL)
		fileName := fmt.Sprintf("avatar_%s.jpg", userID.String())
		seaweedURL, err := u.Seaweed.UploadFromURL(ctx, avatarSource, fileName, "/avatars")
		if err == nil && seaweedURL != "" {
			avatarSource = seaweedURL
		}
		user.DefaultAvatarURL = &avatarSource
	}
	if req.ActiveRole != nil && *req.ActiveRole != "" {
		user.ActiveRole = *req.ActiveRole
	}
	if req.PhoneNumberIndex != nil {
		if *req.PhoneNumberIndex < 0 || *req.PhoneNumberIndex >= len(user.PhoneNumbers) {
			return dto.GetGeneralProfileResponse{}, errors.New(string(i18n.KeyUserPhoneIndexInvalid))
		}
		for i := range user.PhoneNumbers {
			user.PhoneNumbers[i].IsPrimary = i == *req.PhoneNumberIndex
		}
	}
	if err := u.Repo.UpdateUser(ctx, user); err != nil {
		return dto.GetGeneralProfileResponse{}, err
	}
	return u.GetGeneralProfile(ctx, userID)
}
