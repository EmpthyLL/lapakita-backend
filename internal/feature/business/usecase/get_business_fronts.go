package usecase

import (
	"context"
	"errors"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) GetFronts(ctx context.Context, req dto.GetBusinessFrontsRequest) ([]dto.BusinessFrontResponse, api.PaginationMeta, error) {
	businessUUID, err := api.DecodeRequest(req.BusinessID)
	if err != nil {
		return nil, api.PaginationMeta{}, api.ErrInvalidID
	}

	business, err := u.businessRepo.FindByID(ctx, businessUUID)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}
	if business == nil {
		return nil, api.PaginationMeta{}, errors.New("business not found")
	}

	fronts, meta, err := u.businessRepo.FindFrontsByBusinessID(ctx, business.ID, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	res := make([]dto.BusinessFrontResponse, 0, len(fronts))
	for _, f := range fronts {
		res = append(res, dto.BusinessFrontResponse{
			ID:         api.EncodeResponse(f.ID),
			BusinessID: api.EncodeResponse(f.BusinessID),
			Platform:   f.Platform,
			StoreName:  f.StoreName,
			StoreURL:   f.StoreURL,
			IsActive:   f.IsActive,
			CreatedAt:  f.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:  f.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return res, meta, nil
}
