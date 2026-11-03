package dto

import (
	"lapakita-backend/pkg/api"

	"github.com/google/uuid"
)

type ListBusinessRequest struct {
	api.BasePaginationRequest
	UserID         *uuid.UUID `form:"user_id" binding:"omitempty"`
	BusinessTypeID *uuid.UUID `form:"business_type_id" binding:"omitempty"`
	Search         string     `form:"search" binding:"omitempty"`
}

type ListBusinessItemResponse struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	LogoURL          *string `json:"logo_url,omitempty"`
	BusinessTypeName string  `json:"business_type_name"`
	CreatedAt        string  `json:"created_at"`
}
