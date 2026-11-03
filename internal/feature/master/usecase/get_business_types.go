package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
)

func (u *MasterUsecase) GetBusinessTypes(ctx context.Context, lang string, req *dto.GetBusinessTypesRequest) ([]dto.BusinessTypeResponse, api.PaginationMeta, error) {
	activeLang := req.Lang
	if activeLang == "" {
		activeLang = lang
	}
	if activeLang != "id" && activeLang != "en" {
		activeLang = "en"
	}

	items, meta, err := u.Repo.GetBusinessTypes(ctx, activeLang, req)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	responses := make([]dto.BusinessTypeResponse, 0, len(items))
	for _, item := range items {
		label := item.LabelLang.ID
		groupName := item.GroupNameLang.ID
		if activeLang == "en" {
			label = item.LabelLang.EN
			groupName = item.GroupNameLang.EN
		}
		if label == "" {
			label = item.LabelLang.EN
		}
		if groupName == "" {
			groupName = item.GroupNameLang.EN
		}

		recommendedLandmarks, err := json.Marshal(item.RecommendedLandmarks)
		if err != nil {
			return nil, api.PaginationMeta{}, fmt.Errorf("marshal business type recommended landmarks: %w", err)
		}

		responses = append(responses, dto.BusinessTypeResponse{
			ID:                         api.EncodeResponse(item.ID),
			Label:                      label,
			GroupName:                  groupName,
			DefaultBEPMonths:           item.DefaultBEPMonths,
			DefaultCapital:             item.DefaultCapital,
			AvgGrossMarginRatio:        item.AvgGrossMarginRatio,
			IndustryRentToRevenueRatio: item.IndustryRentToRevenueRatio,
			PermanencePresets:          item.PermanencePresets,
			RecommendedLandmarks:       recommendedLandmarks,
		})
	}

	return responses, meta, nil
}
