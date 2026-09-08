package dto

type PersonaProfileResponse struct {
	Role        string `json:"role"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
}

type UpdatePersonaRequest struct {
	DisplayName string `json:"display_name" binding:"required,max=255"`
	AvatarURL   string `json:"avatar_url" binding:"omitempty,url"`
}
