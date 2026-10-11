package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func (u *StallUsecase) Delete(ctx context.Context, ownerID uuid.UUID, id uuid.UUID) error {
	stall, err := u.repo.FindByID(ctx, id)
	if err != nil || stall == nil {
		return errors.New("stall not found")
	}
	if stall.StallOwnerID != ownerID {
		return errors.New("unauthorized delete action")
	}
	return u.repo.Delete(ctx, id)
}
