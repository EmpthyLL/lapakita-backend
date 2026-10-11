package usecase

import (
	"context"
	"errors"

	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func (u *UserUsecase) DeletePhoneNumber(ctx context.Context, userID uuid.UUID, index int) error {
	user, err := u.repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	if index < 0 || index >= len(user.PhoneNumbers) {
		return errors.New(string(i18n.KeyUserPhoneIndexInvalid))
	}
	if user.PhoneNumbers[index].IsPrimary {
		return errors.New(string(i18n.KeyUserPhoneCannotDeletePrimary))
	}
	user.PhoneNumbers = append(user.PhoneNumbers[:index], user.PhoneNumbers[index+1:]...)
	if len(user.PhoneNumbers) > 0 {
		hasPrimary := false
		for _, phone := range user.PhoneNumbers {
			if phone.IsPrimary {
				hasPrimary = true
				break
			}
		}
		if !hasPrimary {
			user.PhoneNumbers[0].IsPrimary = true
		}
	}
	return u.repo.UpdateUser(ctx, user)
}
