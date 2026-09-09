package dto

import "lapakita-backend/pkg/api"

type GetPhoneNumbersRequest struct {
	api.BasePaginationRequest
	Number string `form:"number" binding:"omitempty"`
}

type PhoneNumberItem struct {
	Number    string   `json:"number"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"`
}

type AddPhoneNumberRequest struct {
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles" binding:"omitempty"`
}

type UpdatePhoneNumberRequest struct {
	Number    string   `json:"number" binding:"required,max=32"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles" binding:"omitempty"`
}
