package usecase

import (
	"context"
	"errors"

	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) DeleteFront(ctx context.Context, frontID string) error {
	frontUUID, err := api.DecodeRequest(frontID)
	if err != nil {
		return api.ErrInvalidID
	}

	front, err := u.businessRepo.FindFrontByID(ctx, frontUUID)
	if err != nil {
		return err
	}
	if front == nil {
		return errors.New("business front not found")
	}

	return u.businessRepo.DeleteFront(ctx, front.ID)
}
