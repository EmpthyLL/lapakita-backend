package impl

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/google/uuid"
)

func GetPhoneNumbers(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, req dto.GetPhoneNumbersRequest) ([]dto.PhoneNumberItem, api.PaginationMeta, error) {
	pagedItems, meta, err := u.Repo.GetPhoneNumbers(ctx, userID, &req)
	if err != nil {
		if err.Error() == "user not found" {
			return nil, api.PaginationMeta{}, errors.New(string(i18n.KeyUserNotFound))
		}
		return nil, api.PaginationMeta{}, err
	}
	return pagedItems, meta, nil
}
