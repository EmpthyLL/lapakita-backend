package dto

type SubmitContactInquiryRequest struct {
	Name        string  `json:"name" binding:"required,max=255"`
	Email       string  `json:"email" binding:"required,email,max=255"`
	Whatsapp    *string `json:"whatsapp" binding:"omitempty,max=32"`
	Persona     string  `json:"persona" binding:"required,max=64"`
	InquiryType string  `json:"inquiry_type" binding:"required,max=64"`
	Message     string  `json:"message" binding:"required"`
}

type ContactInquiryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}
