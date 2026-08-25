package dto

type UserPayload struct {
	ID                    string  `json:"id"`
	Name                  string  `json:"name"`
	Email                 string  `json:"email"`
	Phone                 string  `json:"phone,omitempty"`
	AvatarURL             *string `json:"avatarUrl,omitempty"`
	ActiveRole            string  `json:"activeRole"`
	SubscriptionPlan      string  `json:"subscriptionPlan"`
	SubscriptionExpiresAt *string `json:"subscriptionExpiresAt,omitempty"`
	Token                 string  `json:"token"` // Access Token
}

type AuthResponseData struct {
	User         UserPayload `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
