package dto

import "lapakita-backend/pkg/api"

type GetDocumentRequest struct {
	api.BasePaginationRequest
	Name string `form:"name" binding:"omitempty"`
	NIK  string `form:"nik" binding:"omitempty"`
}

type UploadDocumentRequest struct {
	FullNameKTP  string `json:"full_name_ktp" binding:"required,max=255"`
	NIK          string `json:"nik" binding:"required,len=16"`
	DomicileCity string `json:"domicile_city" binding:"required,max=128"`
	KTPPhoto     string `json:"ktp_photo" binding:"required"`
}

type GetDocumentResponse struct {
	ID           string `json:"id"`
	FullNameKTP  string `json:"full_name_ktp"`
	NIK          string `json:"nik"`
	KTPPhotoURL  string `json:"ktp_photo_url"`
	DomicileCity string `json:"domicile_city"`
}
