package dto

import "lapakita-backend/pkg/api"

type GetDocumentRequest struct {
	api.BasePaginationRequest
	Name           string `form:"name" binding:"omitempty"`
	DocumentNumber string `form:"document_number" binding:"omitempty"`
	DocumentType   string `form:"document_type" binding:"omitempty,oneof=national_id passport residence_permit"`
}

type UploadDocumentRequest struct {
	DocumentType     string `json:"document_type" binding:"required,oneof=national_id passport residence_permit"`
	FullNameIdentity string `json:"full_name_identity" binding:"required,max=255"`
	DocumentNumber   string `json:"document_number" binding:"required,max=64"`
	DocumentPhoto    string `json:"document_photo" binding:"required"`
}

type GetDocumentResponse struct {
	ID               string `json:"id"`
	DocumentType     string `json:"document_type"`
	FullNameIdentity string `json:"full_name_identity"`
	DocumentNumber   string `json:"document_number"`
	DocumentPhotoURL string `json:"document_photo_url"`
}
