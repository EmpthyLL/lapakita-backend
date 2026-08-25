package dto

type ResetPasswordRequest struct {
	VerificationToken string `json:"verification_token" binding:"required"`
	NewPassword       string `json:"new_password" binding:"required"`
}
