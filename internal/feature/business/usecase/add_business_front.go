package usecase

import (
	"context"
	"errors"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) AddFront(ctx context.Context, req dto.AddBusinessFrontRequest) (*dto.BusinessFrontResponse, error) {
	businessUUID, err := api.DecodeRequest(req.BusinessID)
	if err != nil {
		return nil, api.ErrInvalidID
	}

	business, err := u.businessRepo.FindByID(ctx, businessUUID)
	if err != nil {
		return nil, err
	}
	if business == nil {
		return nil, errors.New("business not found")
	}

	var storeURL *string
	if req.StoreURL != nil && strings.TrimSpace(*req.StoreURL) != "" {
		cleaned := strings.TrimSpace(*req.StoreURL)
		storeURL = &cleaned
	}

	front := &entity.BusinessFront{
		BusinessID: business.ID,
		Platform:   strings.TrimSpace(req.Platform),
		StoreName:  strings.TrimSpace(req.StoreName),
		StoreURL:   storeURL,
		IsActive:   true,
	}

	if err := u.businessRepo.CreateFront(ctx, front); err != nil {
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
