package dto

import "lapakita-backend/pkg/api"

type GetAreaDetailRequest struct {
	api.BasePaginationRequest
	Search string `form:"search" binding:"required"`
}

type AreaDetailResponseData struct {
	Formatted      string  `json:"formatted"`
	StreetAddress  string  `json:"street_address"` // Nama jalan murni
	Suburb         string  `json:"suburb"`         // Kelurahan / Desa
	District       string  `json:"district"`       // Kecamatan
	City           string  `json:"city"`           // Kota / Kabupaten (Bersih)
	Province       string  `json:"province"`       // Provinsi (Bersih)
	Country        string  `json:"country"`        // Indonesia
	CountryCode    string  `json:"country_code"`   // ID
	PostalCode     string  `json:"postal_code"`    // Kode Pos
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	MapURL         string  `json:"map_url"`
	EmbeddedMapURL string  `json:"embedded_map_url"`
}

// Struct Internal Decode JSON Geoapify
type GeoapifyProperties struct {
	Formatted   string  `json:"formatted"`
	Name        string  `json:"name"`
	Street      string  `json:"street"`
	Suburb      string  `json:"suburb"`
	District    string  `json:"district"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Postcode    string  `json:"postcode"`
	ResultType  string  `json:"result_type"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

type GeoapifyFeature struct {
	Properties GeoapifyProperties `json:"properties"`
}

type GeoapifyResponse struct {
	Features []GeoapifyFeature `json:"features"`
}
