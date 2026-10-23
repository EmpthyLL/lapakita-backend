package impl

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func UpdatePersonaProfile(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, role string, req dto.UpdatePersonaRequest) (dto.PersonaProfileResponse, error) {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.PersonaProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}
	if user.RoleProfiles == nil {
		user.RoleProfiles = entity.RoleProfiles{}
	}
	avatarURL := ""
	if profile, exists := user.RoleProfiles[role]; exists {
		avatarURL = profile.AvatarURL
	}
	if req.AvatarURL != nil && *req.AvatarURL != "" {
		sourceURL := strings.TrimSpace(*req.AvatarURL)
		fileName := fmt.Sprintf("persona_%s_%s.jpg", role, userID.String())
		seaweedURL, uploadErr := u.Seaweed.UploadFromURL(ctx, sourceURL, fileName, "/avatars")
		if uploadErr == nil && seaweedURL != "" {
			avatarURL = seaweedURL
		} else {
			avatarURL = sourceURL
		}
	}
	user.RoleProfiles[role] = entity.RoleProfileItem{DisplayName: req.DisplayName, AvatarURL: avatarURL}
	if req.PhoneNumberIndex != nil {
		targetIdx := *req.PhoneNumberIndex
		if targetIdx < 0 || targetIdx >= len(user.PhoneNumbers) {
			return dto.PersonaProfileResponse{}, errors.New(string(i18n.KeyUserPhoneIndexInvalid))
		}
		user.PhoneNumbers.AssignRoleToPhoneNumberIndex(targetIdx, role)
	}
	if err := u.Repo.UpdateUser(ctx, user); err != nil {
		return dto.PersonaProfileResponse{}, err
	}
	return GetPersonaProfile(ctx, u, userID, role)
}
