package dto

import (
	"lapakita-backend/pkg/api"
)

type GetBusinessFrontsRequest struct {
	api.BasePaginationRequest
	BusinessID string `form:"-" binding:"-"`
	Search     string `form:"search" binding:"omitempty"`
	Platform   string `form:"platform" binding:"omitempty"`
}

type BusinessFrontResponse struct {
	ID         string  `json:"id"`
	BusinessID string  `json:"business_id"`
	Platform   string  `json:"platform"`
	StoreName  string  `json:"store_name"`
	StoreURL   *string `json:"store_url,omitempty"`
	IsActive   bool    `json:"is_active"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
