package dto

type PersonaDetail struct {
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	DialCode    string `json:"dial_code"`
	Phone       string `json:"phone"`
}

type UserPayload struct {
	ID                    string                   `json:"id"`
	DefaultName           string                   `json:"default_name"`
	DefaultAvatarURL      *string                  `json:"default_avatar_url"`
	DefaultPhone          string                   `json:"default_phone"`
	DefaultDialCode       string                   `json:"default_dial_code"`
	Email                 string                   `json:"email"`
	IsPasswordSet         bool                     `json:"is_password_set"`
	ActiveRole            string                   `json:"active_role"`
	SubscriptionPlan      string                   `json:"subscription_plan"`
	SubscriptionExpiresAt *string                  `json:"subscription_expires_at"`
	Personas              map[string]PersonaDetail `json:"personas"`
	Token                 string                   `json:"token"`
}

type PhonePayload struct {
	DialCode  string   `json:"dial_code"`
	Number    string   `json:"number"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type AuthResponseData struct {
	User         UserPayload `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
