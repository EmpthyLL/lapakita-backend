package dto

type GetBusinessByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetBusinessResponse struct {
	ID               string  `json:"id"`
	UserID           string  `json:"user_id"`
	Name             string  `json:"name"`
	LogoURL          *string `json:"logo_url,omitempty"`
	BusinessTypeName string  `json:"business_type_name"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}
