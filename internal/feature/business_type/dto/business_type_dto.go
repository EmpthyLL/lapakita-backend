package dto

import (
	"encoding/json"

	"lapakita-backend/pkg/api"
)

type GetBusinessTypesRequest struct {
	api.BasePaginationRequest
	Search string `form:"search" binding:"omitempty"`
	Group  string `form:"group" binding:"omitempty"`
	Lang   string `form:"lang" binding:"omitempty"`
}

type BusinessTypeResponse struct {
	ID                         string          `json:"id"`
	Label                      string          `json:"label"`
	GroupName                  string          `json:"group_name"`
	DefaultBEPMonths           int             `json:"default_bep_months"`
	DefaultCapital             float64         `json:"default_capital"`
	AvgGrossMarginRatio        float64         `json:"avg_gross_margin_ratio"`
	IndustryRentToRevenueRatio float64         `json:"industry_rent_to_revenue_ratio"`
	PermanencePresets          json.RawMessage `json:"permanence_presets"`
	RecommendedLandmarks       json.RawMessage `json:"recommended_landmarks"`
}
