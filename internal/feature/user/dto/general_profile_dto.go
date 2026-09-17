package dto

type GetGeneralProfileResponse struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	Email            string      `json:"email"`
	DefaultAvatarURL string      `json:"default_avatar_url"`
	ActiveRole       string      `json:"active_role"`
	Phone            PhoneNumber `json:"phone"`
}

type UpdateGeneralProfileRequest struct {
	Name              string  `json:"name" binding:"required,max=255"`
	DefaultAvatarURL  *string `json:"default_avatar_url" binding:"omitempty"`
	PrimaryPhoneIndex *int    `json:"primary_phone_index" binding:"omitempty,min=0"`
	ActiveRole        *string `json:"active_role" binding:"omitempty,oneof=tenant owner supplier"`
}

type PhoneNumber struct {
	Index    int    `json:"index"`
	DialCode string `json:"dial_code"`
	Number   string `json:"number"`
}
