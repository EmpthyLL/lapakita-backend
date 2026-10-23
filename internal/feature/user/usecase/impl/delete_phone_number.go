package impl

import (
	"context"
	"errors"

	usermessages "lapakita-backend/internal/feature/user/messages"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func DeletePhoneNumber(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, index int) error {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	if index < 0 || index >= len(user.PhoneNumbers) {
		return errors.New(string(usermessages.KeyUserPhoneIndexInvalid))
	}
	if user.PhoneNumbers[index].IsPrimary {
		return errors.New(string(usermessages.KeyUserPhoneCannotDeletePrimary))
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
	return u.Repo.UpdateUser(ctx, user)
}
