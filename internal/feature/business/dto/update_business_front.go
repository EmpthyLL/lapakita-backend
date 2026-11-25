package dto

type UpdateBusinessFrontRequest struct {
	ID        string  `json:"-"`
	Platform  string  `json:"platform" binding:"omitempty,oneof=offline_stall shopee tokopedia gofood grabfood website other"`
	StoreName string  `json:"store_name" binding:"omitempty,min=2,max=255"`
	StoreURL  *string `json:"store_url" binding:"omitempty"`
	IsActive  bool    `json:"is_active" binding:"omitempty"`
}
