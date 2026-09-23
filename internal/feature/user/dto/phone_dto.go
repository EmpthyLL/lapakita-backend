package dto

import "lapakita-backend/pkg/api"

type GetPhoneNumbersRequest struct {
	api.BasePaginationRequest
	Search   string `form:"search" binding:"omitempty"`
	DialCode string `form:"dial_code" binding:"omitempty"`
	Role     string `form:"role" binding:"omitempty"`
}

type PhoneNumberItem struct {
	Index     int      `json:"index"`
	Label     string   `json:"label"` // e.g. "Kontak Utama Usaha", "WhatsApp Logistik"
	DialCode  string   `json:"dial_code"`
	Number    string   `json:"number"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type AddPhoneNumberRequest struct {
	Label     string   `json:"label" binding:"required,max=100"`
	DialCode  string   `json:"dial_code" binding:"required,max=8"`
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles" binding:"omitempty"`
}

type UpdatePhoneNumberRequest struct {
	Label     string   `json:"label" binding:"required,max=100"`
	DialCode  string   `json:"dial_code" binding:"required,max=8"`
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles" binding:"omitempty"`
}
