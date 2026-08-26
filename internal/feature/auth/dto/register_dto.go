package dto

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=2"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Phone    string `json:"phone" binding:"required"`
}

type CompleteProfileRequest struct {
	SetupToken string `json:"setup_token" binding:"required"`
	Name       string `json:"name" binding:"required,min=2"`
	Password   string `json:"password" binding:"required,min=6"`
	Phone      string `json:"phone" binding:"required"`
	AvatarURL  string `json:"avatar_url" binding:"omitempty,url"`
}

type GoogleSetupPresetResponse struct {
	Email              string `json:"email"`
	Name               string `json:"name"`
	AvatarURL          string `json:"avatar_url"`
	SetupToken         string `json:"setup_token"`
	IsProfileCompleted bool   `json:"is_profile_completed"`
}
