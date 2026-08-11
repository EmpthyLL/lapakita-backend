package dto

type GetAreaRequest struct {
	Categories string  `form:"categories" binding:"required"`
	Lat        float64 `form:"lat" binding:"required"`
	Lon        float64 `form:"lon" binding:"required"`
	Radius     int     `form:"radius" binding:"required,min=100,max=10000"`
}

type AreaResponse struct {
	Name           string  `json:"name"`
	Country        string  `json:"country"`
	City           string  `json:"city"`
	Street         string  `json:"street"`
	Lon            float64 `json:"lon"`
	Lat            float64 `json:"lat"`
	Formatted      string  `json:"formatted"`
	AreaID         string  `json:"area_id"`
	GoogleMapURL   string  `json:"google_map_url"`
	GoogleEmbedURL string  `json:"google_embed_url"`
}
