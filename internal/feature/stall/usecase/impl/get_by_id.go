package impl

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/stall/dto"
	usecase "lapakita-backend/internal/feature/stall/usecase"

	"github.com/google/uuid"
)

func GetByID(ctx context.Context, u *usecase.StallUsecase, id uuid.UUID) (dto.StallDetailResponse, error) {
	stall, err := u.StallRepository.FindByID(ctx, id)
	if err != nil || stall == nil {
		return dto.StallDetailResponse{}, errors.New("stall not found")
	}
	ownerUser, _ := u.AuthRepository.FindUserByID(ctx, stall.StallOwnerID)
	return usecase.MapToDetailResponse(stall, ownerUser), nil
}
