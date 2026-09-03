package dto

import (
	"lapakita-backend/pkg/api"

	"github.com/google/uuid"
)

type GetSimilarStallsRequest struct {
	api.BasePaginationRequest
	ID uuid.UUID `uri:"id" binding:"required"`
}
