package dto

import (
	"lapakita-backend/pkg/api"

	"github.com/google/uuid"
)

type ListBusinessRequest struct {
	api.BasePaginationRequest
	UserID         uuid.UUID  `form:"-" binding:"-"`
	BusinessTypeID *uuid.UUID `form:"business_type_id" binding:"omitempty"`
	Search         string     `form:"search" binding:"omitempty"`
}

type BusinessFrontItemResponse struct {
	ID        string  `json:"id"`
	Platform  string  `json:"platform"`
	StoreName string  `json:"store_name"`
	StoreURL  *string `json:"store_url,omitempty"`
	IsActive  bool    `json:"is_active"`
}

type ListBusinessItemResponse struct {
	ID               string                      `json:"id"`
	Name             string                      `json:"name"`
	LogoURL          *string                     `json:"logo_url,omitempty"`
	BusinessTypeName string                      `json:"business_type_name"`
	Fronts           []BusinessFrontItemResponse `json:"fronts"`
	TotalFronts      int64                       `json:"total_fronts"`
	CreatedAt        string                      `json:"created_at"`
}
