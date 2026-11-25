package usecase

import (
	"context"
	"errors"
	"strings"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) UpdateFront(ctx context.Context, req dto.UpdateBusinessFrontRequest) (*dto.BusinessFrontResponse, error) {
	frontUUID, err := api.DecodeRequest(req.ID)
	if err != nil {
		return nil, api.ErrInvalidID
	}

	front, err := u.businessRepo.FindFrontByID(ctx, frontUUID)
	if err != nil {
		return nil, err
	}
	if front == nil {
		return nil, errors.New("business front not found")
	}

	if req.StoreURL != nil {
		cleaned := strings.TrimSpace(*req.StoreURL)
		if cleaned == "" {
			front.StoreURL = nil
		} else {
			front.StoreURL = &cleaned
		}
	}

	if err := u.businessRepo.UpdateFront(ctx, front); err != nil {
		return nil, err
	}

	return &dto.BusinessFrontResponse{
		ID:         api.EncodeResponse(front.ID),
		BusinessID: api.EncodeResponse(front.BusinessID),
		Platform:   front.Platform,
		StoreName:  front.StoreName,
		StoreURL:   front.StoreURL,
		IsActive:   front.IsActive,
		CreatedAt:  front.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:  front.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
