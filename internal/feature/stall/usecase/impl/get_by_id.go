package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/stall/dto"
	sharedusecase "lapakita-backend/internal/feature/stall/usecase"

	"github.com/google/uuid"
)

func (u *StallUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.StallDetailResponse, error) {
	stall, err := u.repo.FindByID(ctx, id)
	if err != nil || stall == nil {
		return dto.StallDetailResponse{}, errors.New("stall not found")
	}
	ownerUser, _ := u.userRepo.FindUserByID(ctx, stall.StallOwnerID)
	return sharedusecase.MapToDetailResponse(stall, ownerUser), nil
}
