package dto

type GetGeneralProfileResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	DefaultAvatarURL string `json:"default_avatar_url"`
}

type UpdateGeneralProfileRequest struct {
	Name             string `json:"name" binding:"required,max=255"`
	DefaultAvatarURL string `json:"default_avatar_url" binding:"omitempty,url"`
}
