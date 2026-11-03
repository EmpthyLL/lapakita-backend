package dto

import (
	"github.com/google/uuid"
)

type CreateBusinessRequest struct {
	UserID         uuid.UUID `json:"user_id" binding:"required"`
	BusinessTypeID string    `json:"business_type_id" binding:"required"`
	Name           string    `json:"name" binding:"required,min=2,max=255"`
	LogoURL        *string   `json:"logo_url" binding:"omitempty,url"`
}

type CreateBusinessResponse struct {
	ID             string  `json:"id"`
	UserID         string  `json:"user_id"`
	BusinessTypeID string  `json:"business_type_id"`
	Name           string  `json:"name"`
	LogoURL        *string `json:"logo_url,omitempty"`
	CreatedAt      string  `json:"created_at"`
}
