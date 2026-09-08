package dto

type GetGeneralProfileResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	DefaultAvatarURL string `json:"default_avatar_url"`
	PrimaryPhone     string `json:"primary_phone"`
}

type UpdateGeneralProfileRequest struct {
	Name             string  `json:"name" binding:"required,max=255"`
	DefaultAvatarURL *string `json:"default_avatar_url" binding:"omitempty"`
	PhoneNumber      *string `json:"phone_number" binding:"omitempty,max=32"`
}
