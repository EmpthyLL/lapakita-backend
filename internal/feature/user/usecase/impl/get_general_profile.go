package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func (u *UserUsecase) GetGeneralProfile(ctx context.Context, userID uuid.UUID) (dto.GetGeneralProfileResponse, error) {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return dto.GetGeneralProfileResponse{}, errors.New(string(i18n.KeyUserNotFound))
	}
	avatar := ""
	if user.DefaultAvatarURL != nil {
		avatar = *user.DefaultAvatarURL
	}
	return dto.GetGeneralProfileResponse{
		Name: user.Name, Email: user.Email, DefaultAvatarURL: avatar, ActiveRole: user.ActiveRole,
		Phone: dto.PhoneNumber{Index: user.PhoneNumbers.GetPrimaryIndex(), DialCode: user.PhoneNumbers.GetPrimaryDialCode(), Number: user.PhoneNumbers.GetPrimaryNumber()},
	}, nil
}
