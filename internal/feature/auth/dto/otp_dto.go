package dto

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	Mode  string `json:"mode" binding:"required,oneof=register reset_password"`
}

type VerifyOTPRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Mode    string `json:"mode" binding:"required,oneof=register reset_password"`
	OTPCode string `json:"otp_code" binding:"required,len=6"`
}

type OTPStatePayload struct {
	Email string `json:"email"`
	Mode  string `json:"mode"`
}

type VerifyOTPResponse struct {
	VerificationToken string            `json:"verification_token,omitempty"`
	AuthData          *AuthResponseData `json:"auth_data,omitempty"`
}
