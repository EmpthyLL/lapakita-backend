package dto

// SaveHistoryRequest digunakan saat pengguna menekan/pilih satu item area
type SaveHistoryRequest struct {
	AreaGeneralResponseData
}

// AreaHistoryItemResponse membungkus data area beserta timestamp pencariannya
type AreaHistoryItemResponse struct {
	AreaGeneralResponseData
	SearchedAt string `json:"searched_at"` // Format RFC3339 (e.g. "2026-08-30T23:54:25Z")
}

// DeleteHistoryItemRequest digunakan untuk menghapus 1 item spesifik berdasarkan FullLabel
type DeleteHistoryItemRequest struct {
	FullLabel string `json:"full_label" binding:"required"`
}
