package dto

type GoogleAuthRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}

type CompleteProfileRequest struct {
	Name      string `json:"name" binding:"required,min=2"`
	DialCode  string `json:"dial_code" binding:"required,max=8"`
	Phone     string `json:"phone" binding:"required"`
	AvatarURL string `json:"avatar_url" binding:"omitempty,url"`
}
