package dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type CompleteProfileRequest struct {
	SetupToken string `json:"setup_token" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	AvatarURL  string `json:"avatar_url"`
}

type GoogleSetupPresetResponse struct {
	Email              string `json:"email"`
	Name               string `json:"name"`
	AvatarURL          string `json:"avatar_url"`
	SetupToken         string `json:"setup_token"`
	IsProfileCompleted bool   `json:"is_profile_completed"`
}
