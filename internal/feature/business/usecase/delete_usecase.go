package usecase

import (
	"context"
	"errors"

	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) Delete(ctx context.Context, id string) error {
	businessUUID, err := api.DecodeRequest(id)
	if err != nil {
		return api.ErrInvalidID
	}

	business, err := u.businessRepo.FindByID(ctx, businessUUID)
	if err != nil {
		return err
	}
	if business == nil {
		return errors.New("business not found")
	}

	return u.businessRepo.Delete(ctx, businessUUID)
}
