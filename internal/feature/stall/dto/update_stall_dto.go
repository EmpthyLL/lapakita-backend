package dto

type UpdateStallRequest struct {
	ID string `json:"id" binding:"required"`
	CreateStallRequest
}
