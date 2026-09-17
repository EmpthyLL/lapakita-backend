package dto

import "lapakita-backend/pkg/api"

type GetPhoneNumbersRequest struct {
	api.BasePaginationRequest
	Search string `form:"search" binding:"omitempty"`
}

type PhoneNumberItem struct {
	Index     int      `json:"index"`
	DialCode  string   `json:"dial_code"`
	Number    string   `json:"number"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type AddPhoneNumberRequest struct {
	DialCode  string   `json:"dial_code" binding:"required,max=8"`
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles" binding:"omitempty"`
}

type UpdatePhoneNumberRequest struct {
	DialCode  string   `json:"dial_code" binding:"required,max=8"`
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles" binding:"omitempty"`
}
