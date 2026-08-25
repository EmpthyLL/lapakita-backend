package dto

type SendOTPRequest struct {
	Email string `json:"email" binding:"required"`
	Mode  string `json:"mode" binding:"required"`
}

type VerifyOTPRequest struct {
	StatePayload string `json:"state_payload" binding:"required"`
	OTPCode      string `json:"otp_code" binding:"required"`
}

type OTPStatePayload struct {
	Email string `json:"email"`
	Mode  string `json:"mode"`
}
