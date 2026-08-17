package dto

import "lapakita-backend/pkg/api"

type GetAreaGeneralRequest struct {
	api.BasePaginationRequest
	Search string `form:"search" binding:"required"`
}

type AreaGeneralResponseData struct {
	Type        string `json:"type"`         // "country", "province", "city", "district", "street"
	Title       string `json:"title"`        // Contoh: "Jalan Delimas" / "Surabaya"
	Subtitle    string `json:"subtitle"`     // Contoh: "Surabaya, Jawa Timur, Indonesia"
	FullLabel   string `json:"full_label"`   // Contoh: "Jalan Delimas, Surabaya, Jawa Timur, Indonesia"
	Country     string `json:"country"`      // Contoh: "Indonesia"
	CountryCode string `json:"country_code"` // Contoh: "ID"
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
}

type GetAreaGeneralResponse = api.PaginatedResponse[[]AreaGeneralResponseData]
