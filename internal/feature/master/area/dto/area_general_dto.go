package dto

import "lapakita-backend/pkg/api"

type GetAreaGeneralRequest struct {
	api.BasePaginationRequest
	Search string `form:"search" binding:"omitempty"`
}

type AreaGeneralResponseData struct {
	Type        string `json:"type"`               // "country", "province", "city", "district", "suburb", "street"
	Title       string `json:"title"`              // Contoh: "Jalan Delimas" / "Tebet" / "Surabaya"
	Subtitle    string `json:"subtitle"`           // Contoh: "Surabaya, Jawa Timur, Indonesia"
	FullLabel   string `json:"full_label"`         // Contoh: "Jalan Delimas, Tebet, Surabaya, Jawa Timur, Indonesia"
	Country     string `json:"country"`            // Contoh: "Indonesia"
	CountryCode string `json:"country_code"`       // Contoh: "ID"
	City        string `json:"city,omitempty"`     // Kota / Kabupaten (Clean)
	Province    string `json:"province,omitempty"` // Provinsi (Clean)
	District    string `json:"district,omitempty"` // Kecamatan
	Suburb      string `json:"suburb,omitempty"`   // Kelurahan / Desa
}
