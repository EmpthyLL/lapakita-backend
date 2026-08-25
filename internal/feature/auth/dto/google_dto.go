package dto

type GoogleAuthRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}

type GoogleTokenPayload struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
