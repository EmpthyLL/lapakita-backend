package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func (u *UserUsecase) UpdatePhoneNumber(ctx context.Context, userID uuid.UUID, index int, req dto.UpdatePhoneNumberRequest) error {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	if index < 0 || index >= len(user.PhoneNumbers) {
		return errors.New(string(i18n.KeyUserPhoneIndexInvalid))
	}
	for i, phone := range user.PhoneNumbers {
		if i != index && phone.DialCode == req.DialCode && phone.Number == req.Number {
			return errors.New(string(i18n.KeyUserPhoneDuplicate))
		}
	}
	isCurrentlyPrimary := user.PhoneNumbers[index].IsPrimary
	if isCurrentlyPrimary && !req.IsPrimary {
		hasOtherPrimary := false
		for i, phone := range user.PhoneNumbers {
			if i != index && phone.IsPrimary {
				hasOtherPrimary = true
				break
			}
		}
		if !hasOtherPrimary {
			return errors.New(string(i18n.KeyUserPhonePrimaryRequired))
		}
	}
	user.PhoneNumbers = TransferRolesToTarget(user.PhoneNumbers, index, req.Roles)
	user.PhoneNumbers[index].Label = req.Label
	user.PhoneNumbers[index].Number = req.Number
	user.PhoneNumbers[index].DialCode = req.DialCode
	user.PhoneNumbers[index].Roles = req.Roles
	if req.IsPrimary {
		for i := range user.PhoneNumbers {
			user.PhoneNumbers[i].IsPrimary = i == index
		}
	}
	return u.Repo.UpdateUser(ctx, user)
}
