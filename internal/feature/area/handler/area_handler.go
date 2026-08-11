package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/internal/feature/area/usecase"
	"lapakita-backend/pkg/api"

	"github.com/gin-gonic/gin"
)

type AreaHandler struct {
	usecase usecase.AreaUsecase
}

func NewAreaHandler(usecase usecase.AreaUsecase) *AreaHandler {
	return &AreaHandler{
		usecase: usecase,
	}
}

func (h *AreaHandler) GetNearbyPlaces(c *gin.Context) {
	var req dto.GetAreaRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		api.ErrorResponse(c, http.StatusBadRequest, "Parameter query tidak valid", err.Error())
		return
	}

	places, err := h.usecase.GetNearbyPlaces(c.Request.Context(), req)
	if err != nil {
		api.ErrorResponse(c, http.StatusInternalServerError, "Gagal mendapatkan data area lokasi", err.Error())
		return
	}

	api.SuccessResponse(c, http.StatusOK, "Berhasil mendapatkan data tempat terdekat", places)
}
