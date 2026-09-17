package dto

import "lapakita-backend/pkg/api"

type GetPhoneNumbersRequest struct {
	api.BasePaginationRequest
	DialCode string `form:"dial_code" binding:"omitempty,max=8"`
	Number   string `form:"number" binding:"omitempty"`
}

type PhoneNumberItem struct {
	Index     int      `json:"phone_number_index"`
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
