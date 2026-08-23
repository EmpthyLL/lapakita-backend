package usecase

import (
	"context"
	"encoding/json"
	"math"

	"lapakita-backend/internal/feature/business_type/dto"
	"lapakita-backend/internal/feature/business_type/repository"
	"lapakita-backend/pkg/api"
)

type BusinessTypeUsecase struct {
	repo *repository.BusinessTypeRepository
}

func NewBusinessTypeUsecase(repo *repository.BusinessTypeRepository) *BusinessTypeUsecase {
	return &BusinessTypeUsecase{repo: repo}
}

func (u *BusinessTypeUsecase) GetBusinessTypes(ctx context.Context, lang string, req *dto.GetBusinessTypesRequest) ([]dto.BusinessTypeResponse, api.PaginationMeta, error) {
	activeLang := req.Lang
	if activeLang == "" {
		activeLang = lang
	}
	if activeLang != "id" && activeLang != "en" {
		activeLang = "en"
	}

	items, total, err := u.repo.GetBusinessTypes(ctx, activeLang, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}

	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	var responses []dto.BusinessTypeResponse
	for _, item := range items {
		labelMap := make(map[string]string)
		groupMap := make(map[string]string)

		_ = json.Unmarshal(item.LabelLang, &labelMap)
		_ = json.Unmarshal(item.GroupNameLang, &groupMap)

		label := labelMap[activeLang]
		if label == "" {
			label = labelMap["en"]
		}

		groupName := groupMap[activeLang]
		if groupName == "" {
			groupName = groupMap["en"]
		}

		responses = append(responses, dto.BusinessTypeResponse{
			ID:                         item.ID.String(),
			Label:                      label,
			GroupName:                  groupName,
			DefaultBEPMonths:           item.DefaultBEPMonths,
			DefaultCapital:             item.DefaultCapital,
			AvgGrossMarginRatio:        item.AvgGrossMarginRatio,
			IndustryRentToRevenueRatio: item.IndustryRentToRevenueRatio,
			PermanencePresets:          item.PermanencePresets,
			RecommendedLandmarks:       item.RecommendedLandmarks,
		})
	}

	meta := api.PaginationMeta{
		TotalItems:  int(total),
		TotalPages:  totalPages,
		CurrentPage: page,
		PerPage:     limit,
		HasNextPage: page < totalPages,
		HasPrevPage: page > 1,
	}

	return responses, meta, nil
}
