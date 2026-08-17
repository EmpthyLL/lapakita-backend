package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/internal/feature/area/usecase"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

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

// GET /api/v1/areas (Search Bar UI Dropdown)
func (h *AreaHandler) SearchGeneral(c *gin.Context) {
	var req dto.GetAreaGeneralRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid))
		return
	}

	areas, meta, err := h.usecase.SearchGeneral(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), areas, meta)
}

// GET /api/v1/areas/detail (Auto-fill Form Input Owner)
func (h *AreaHandler) SearchDetail(c *gin.Context) {
	var req dto.GetAreaDetailRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid))
		return
	}

	details, meta, err := h.usecase.SearchDetail(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), details, meta)
}
