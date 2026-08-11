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

func (h *AreaHandler) SearchArea(c *gin.Context) {
	var req dto.GetAreaRequest

	lang := c.GetHeader("lang")
	if lang == "" {
		lang = "en"
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		msg := "Query parameter invalid"
		if lang == "id" {
			msg = "Parameter query tidak valid"
		}
		api.ErrorResponse(c, http.StatusBadRequest, msg, err.Error())
		return
	}

	areas, err := h.usecase.SearchArea(c.Request.Context(), req, lang)
	if err != nil {
		msg := "Failed to search area"
		if lang == "id" {
			msg = "Gagal mencari data area"
		}
		api.ErrorResponse(c, http.StatusInternalServerError, msg, err.Error())
		return
	}

	successMsg := "Successfully retrieved area autocomplete results"
	if lang == "id" {
		successMsg = "Berhasil mendapatkan hasil pencarian area"
	}

	api.SuccessResponse(c, http.StatusOK, successMsg, areas)
}
