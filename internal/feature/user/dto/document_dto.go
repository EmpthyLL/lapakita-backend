package dto

import "lapakita-backend/pkg/api"

type GetDocumentRequest struct {
	api.BasePaginationRequest
	Search       string `form:"search" binding:"omitempty"`
	DocumentType string `form:"document_type" binding:"omitempty,oneof=national_id passport residence_permit"`
	CountryCode  string `form:"country_code" binding:"omitempty"`
}

type UploadDocumentRequest struct {
	CountryCode      string `json:"country_code" binding:"omitempty,max=8"` // Default: "ID"
	DocumentType     string `json:"document_type" binding:"required,oneof=national_id passport residence_permit"`
	DocumentLabel    string `json:"document_label" binding:"required,max=100"` // e.g. "KTP Utama", "Paspor Indonesia"
	FullNameIdentity string `json:"full_name_identity" binding:"required,max=255"`
	DocumentNumber   string `json:"document_number" binding:"required,max=64"`
	DocumentPhoto    string `json:"document_photo" binding:"required"` // Base64 String
}

type GetDocumentResponse struct {
	ID               string `json:"id"`
	CountryCode      string `json:"country_code"`
	DocumentType     string `json:"document_type"`
	DocumentLabel    string `json:"document_label"`
	FullNameIdentity string `json:"full_name_identity"`
	DocumentNumber   string `json:"document_number"`
	DocumentPhotoURL string `json:"document_photo_url"`
}
