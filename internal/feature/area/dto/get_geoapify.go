package dto

import "lapakita-backend/pkg/api"

// Request dari Frontend (meng-embed BasePaginationRequest)
type GetAreaRequest struct {
	api.BasePaginationRequest
	Search string `form:"search" binding:"required"`
}

// Item Data Respon untuk Frontend
type AreaResponseData struct {
	Formatted      string `json:"formatted"`
	AddressLine1   string `json:"address_line1"`
	AddressLine2   string `json:"address_line2"`
	GoogleMapURL   string `json:"google_map_url"`
	GoogleEmbedURL string `json:"google_embed_url"`
}

// Alias Type Response terstruktur menggunakan Generic
type GetAreaResponse = api.PaginatedResponse[[]AreaResponseData]

// --- Struct Internal khusus decode Geoapify API ---

type GeoapifyProperties struct {
	Formatted    string  `json:"formatted"`
	AddressLine1 string  `json:"address_line1"`
	AddressLine2 string  `json:"address_line2"`
	Lat          float64 `json:"lat"`
	Lon          float64 `json:"lon"`
}

type GeoapifyFeature struct {
	Properties GeoapifyProperties `json:"properties"`
}

type GeoapifyResponse struct {
	Features []GeoapifyFeature `json:"features"`
}
