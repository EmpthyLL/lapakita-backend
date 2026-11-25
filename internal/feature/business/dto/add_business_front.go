package dto

import "github.com/google/uuid"

type AddBusinessFrontRequest struct {
	BusinessID string    `json:"-"`
	UserID     uuid.UUID `json:"-"`
	Platform   string    `json:"platform" binding:"required,oneof=offline_stall shopee tokopedia gofood grabfood website other"`
	StoreName  string    `json:"store_name" binding:"required,min=2,max=255"`
	StoreURL   *string   `json:"store_url" binding:"omitempty,url"`
	IsActive   bool      `json:"is_active" binding:"omitempty"`
}
