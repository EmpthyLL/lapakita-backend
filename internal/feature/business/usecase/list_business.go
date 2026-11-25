package usecase

import (
	"context"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/pkg/api"
)

func (u *BusinessUsecase) List(ctx context.Context, req dto.ListBusinessRequest, lang string) ([]dto.ListBusinessItemResponse, api.PaginationMeta, error) {
	businesses, frontCounts, meta, err := u.businessRepo.List(ctx, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	items := make([]dto.ListBusinessItemResponse, 0, len(businesses))
	for _, b := range businesses {
		var btName string
		if b.BusinessType != nil {
			btName = b.BusinessType.LabelLang.ID
			if lang == "en" {
				btName = b.BusinessType.LabelLang.EN
			}
		}

		fronts := make([]dto.BusinessFrontItemResponse, 0, len(b.Fronts))
		for _, f := range b.Fronts {
			fronts = append(fronts, dto.BusinessFrontItemResponse{
				ID:        api.EncodeResponse(f.ID),
				Platform:  f.Platform,
				StoreName: f.StoreName,
				StoreURL:  f.StoreURL,
				IsActive:  f.IsActive,
			})
		}

		item := dto.ListBusinessItemResponse{
			ID:               api.EncodeResponse(b.ID),
			Name:             b.Name,
			LogoURL:          b.LogoURL,
			BusinessTypeName: btName,
			Fronts:           fronts,
			TotalFronts:      frontCounts[b.ID],
			CreatedAt:        b.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		items = append(items, item)
	}

	return items, meta, nil
}
