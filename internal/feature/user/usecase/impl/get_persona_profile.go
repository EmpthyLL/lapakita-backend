package impl

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func GetPersonaProfile(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, role string) (dto.PersonaProfileResponse, error) {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.PersonaProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}
	phoneIdx, dialCode, number := user.PhoneNumbers.GetPhoneDetailForRole(role)
	profile, exists := user.RoleProfiles[role]
	if !exists {
		defaultAvatar := ""
		if user.DefaultAvatarURL != nil {
			defaultAvatar = *user.DefaultAvatarURL
		}
		return dto.PersonaProfileResponse{Role: role, DisplayName: user.Name, AvatarURL: defaultAvatar, Phone: dto.PhoneNumber{Index: phoneIdx, DialCode: dialCode, Number: number}}, nil
	}
	return dto.PersonaProfileResponse{Role: role, DisplayName: profile.DisplayName, AvatarURL: profile.AvatarURL, Phone: dto.PhoneNumber{Index: phoneIdx, DialCode: dialCode, Number: number}}, nil
}
