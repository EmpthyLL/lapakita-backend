package dto

type PersonaDetail struct {
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	Phone       string `json:"phone"`
}

type UserPayload struct {
	ID                    string                   `json:"id"`
	DefaultName           string                   `json:"default_name"`
	DefaultAvatarURL      *string                  `json:"default_avatar_url"`
	DefaultPhone          string                   `json:"default_phone"`
	Email                 string                   `json:"email"`
	ActiveRole            string                   `json:"active_role"`
	SubscriptionPlan      string                   `json:"subscription_plan"`
	SubscriptionExpiresAt *string                  `json:"subscription_expires_at"`
	PhoneNumbers          []PhonePayload           `json:"phone_numbers"`
	Personas              map[string]PersonaDetail `json:"personas"`
	Token                 string                   `json:"token"`
}

type PhonePayload struct {
	Number    string   `json:"number"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type AuthResponseData struct {
	User         UserPayload `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
