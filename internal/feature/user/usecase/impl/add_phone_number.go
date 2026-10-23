package impl

import (
	"context"
	"errors"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func AddPhoneNumber(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, req dto.AddPhoneNumberRequest) error {
	user, err := u.Repo.FindByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New(string(i18n.KeyUserNotFound))
	}
	for _, phone := range user.PhoneNumbers {
		if phone.DialCode == req.DialCode && phone.Number == req.Number {
			return errors.New(string(i18n.KeyUserPhoneDuplicate))
		}
	}
	user.PhoneNumbers = usecase.TransferRolesToTarget(user.PhoneNumbers, -1, req.Roles)
	newItem := entity.PhoneNumberItem{Label: req.Label, DialCode: req.DialCode, Number: req.Number, IsPrimary: req.IsPrimary, Roles: req.Roles}
	if req.IsPrimary || len(user.PhoneNumbers) == 0 {
		newItem.IsPrimary = true
		for i := range user.PhoneNumbers {
			user.PhoneNumbers[i].IsPrimary = false
		}
	}
	user.PhoneNumbers = append(user.PhoneNumbers, newItem)
	hasPrimary := false
	for _, phone := range user.PhoneNumbers {
		if phone.IsPrimary {
			hasPrimary = true
			break
		}
	}
	if !hasPrimary && len(user.PhoneNumbers) > 0 {
		user.PhoneNumbers[0].IsPrimary = true
	}
	return u.Repo.UpdateUser(ctx, user)
}
