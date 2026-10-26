package usecase

import (
	"context"
	"encoding/json"

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
			ID:                         api.EncodeResponse(item.ID),
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

	return responses, meta, nil
}
