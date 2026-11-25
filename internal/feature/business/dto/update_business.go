package dto

type UpdateBusinessRequest struct {
	ID             string  `uri:"id" binding:"required"`
	BusinessTypeID *string `json:"business_type_id" binding:"omitempty"`
	Name           *string `json:"name" binding:"omitempty,min=2,max=255"`
	LogoURL        *string `json:"logo_url" binding:"omitempty"`
}
