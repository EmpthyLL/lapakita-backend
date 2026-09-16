package dto

type GetGeneralProfileResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	DefaultAvatarURL string `json:"default_avatar_url"`
	ActiveRole       string `json:"active_role"`
	PrimaryDialCode  string `json:"primary_dial_code"`
	PrimaryPhone     string `json:"primary_phone"`
}

type UpdateGeneralProfileRequest struct {
	Name             string  `json:"name" binding:"required,max=255"`
	DefaultAvatarURL *string `json:"default_avatar_url" binding:"omitempty"`
	DialCode         *string `json:"dial_code" binding:"omitempty,max=8"`
	PhoneNumber      *string `json:"phone_number" binding:"omitempty,max=32"`
	ActiveRole       *string `json:"active_role" binding:"omitempty,oneof=tenant owner supplier"`
}
