package dto

// Request dari Frontend
type GetAreaRequest struct {
	Text string `form:"text" binding:"required"`
}

// Respon khusus ke Frontend (Ringkas & UI-Friendly)
type AreaResponse struct {
	Formatted      string `json:"formatted"`        // Alamat lengkap untuk ditampilkan di opsi dropdown
	AddressLine1   string `json:"address_line1"`    // Nama Jalan / Tempat
	AddressLine2   string `json:"address_line2"`    // Kota, Provinsi, Kode Pos
	GoogleMapURL   string `json:"google_map_url"`   // Link ke app/web Google Maps
	GoogleEmbedURL string `json:"google_embed_url"` // Link untuk iframe embed map
}

// Struct internal khusus memetakan JSON dari Geoapify API
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
