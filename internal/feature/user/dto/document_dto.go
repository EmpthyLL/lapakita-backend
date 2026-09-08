package dto

type UploadDocumentRequest struct {
	FullNameKTP  string `json:"full_name_ktp" binding:"required,max=255"`
	NIK          string `json:"nik" binding:"required,len=16"`
	DomicileCity string `json:"domicile_city" binding:"required,max=128"`
	PurposeType  string `json:"purpose_type" binding:"required,oneof=stall_verification supplier_verification lease_contract"`
	KTPPhoto     string `json:"ktp_photo" binding:"required"`
}

type GetDocumentResponse struct {
	FullNameKTP  string `json:"full_name_ktp"`
	NIK          string `json:"nik"`
	KTPPhotoURL  string `json:"ktp_photo_url"`
	DomicileCity string `json:"domicile_city"`
}
