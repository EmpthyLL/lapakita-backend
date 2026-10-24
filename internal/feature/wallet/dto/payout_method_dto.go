package dto

import (
	"time"

	"lapakita-backend/pkg/api"
)

type GetPayoutMethodsRequest struct {
	api.BasePaginationRequest
	Search     string `form:"search" binding:"omitempty"`
	PayoutType string `form:"payout_type" binding:"omitempty,oneof=bank_account e_wallet"`
	Provider   string `form:"provider" binding:"omitempty"`
	IsPrimary  *bool  `form:"is_primary" binding:"omitempty"`
}

type CreatePayoutMethodRequest struct {
	PayoutType        string `json:"payout_type" binding:"required,oneof=bank_account e_wallet"`
	ProviderCode      string `json:"provider_code" binding:"required,max=32"`
	ProviderName      string `json:"provider_name" binding:"required,max=128"`
	AccountNumber     string `json:"account_number" binding:"required,max=64"`
	AccountHolderName string `json:"account_holder_name" binding:"required,max=255"`
}

type PayoutMethodResponse struct {
	ID                string    `json:"id"`
	PayoutType        string    `json:"payout_type"`
	ProviderCode      string    `json:"provider_code"`
	ProviderName      string    `json:"provider_name"`
	AccountNumber     string    `json:"account_number"`
	AccountHolderName string    `json:"account_holder_name"`
	IsPrimary         bool      `json:"is_primary"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
