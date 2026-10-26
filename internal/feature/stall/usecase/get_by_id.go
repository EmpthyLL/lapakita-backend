package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/stall/dto"

	"github.com/google/uuid"
)

func (u *StallUsecase) GetByID(ctx context.Context, id uuid.UUID) (dto.StallDetailResponse, error) {
	stall, err := u.StallRepository.FindByID(ctx, id)
	if err != nil || stall == nil {
		return dto.StallDetailResponse{}, errors.New("stall not found")
	}
	ownerUser, _ := u.AuthRepository.FindUserByID(ctx, stall.StallOwnerID)
	return MapToDetailResponse(stall, ownerUser), nil
}
