package impl

import (
	"context"
	"errors"

	usecase "lapakita-backend/internal/feature/stall/usecase"

	"github.com/google/uuid"
)

func Delete(ctx context.Context, u *usecase.StallUsecase, ownerID uuid.UUID, id uuid.UUID) error {
	stall, err := u.StallRepository.FindByID(ctx, id)
	if err != nil || stall == nil {
		return errors.New("stall not found")
	}
	if stall.StallOwnerID != ownerID {
		return errors.New("unauthorized delete action")
	}
	return u.StallRepository.Delete(ctx, id)
}
