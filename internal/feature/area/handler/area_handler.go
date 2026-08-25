package handler

import (
	"log"
	"net/http"

	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/internal/feature/area/usecase"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type AreaHandler struct {
	usecase *usecase.AreaUsecase // Pastikan tipe pointer murni
}

func NewAreaHandler(usecase *usecase.AreaUsecase) *AreaHandler {
	return &AreaHandler{
		usecase: usecase,
	}
}

// GET /api/v1/areas
func (h *AreaHandler) SearchGeneral(c *gin.Context) {
	// DEBUG LOG 1
	log.Println("[DEBUG] AreaHandler.SearchGeneral dipanggil")

	if h.usecase == nil {
		log.Println("[CRITICAL ERROR] h.usecase bernilai NIL! Wire injection gagal.")
		api.Error(c, http.StatusInternalServerError, "Internal Server Error: Usecase nil")
		return
	}

	var req dto.GetAreaGeneralRequest

	// Gunakan ShouldBindQuery tanpa memblokir jika page/limit tidak dikirim
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("[DEBUG] BindQuery Warn/Err: %v", err)
	}

	// Fallback pembacaan manual jika binding Gin gagal membaca query "search"
	if req.Search == "" {
		req.Search = c.Query("search")
	}

	log.Printf("[DEBUG] Search Query Extracted: '%s'", req.Search)

	areas, meta, err := h.usecase.SearchGeneral(c.Request.Context(), req)
	if err != nil {
		log.Printf("[ERROR] Usecase SearchGeneral Err: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}

	log.Println("[DEBUG] Berhasil dapat data, mengirim response...")
	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), areas, meta)
}

// GET /api/v1/areas/detail
func (h *AreaHandler) SearchDetail(c *gin.Context) {
	log.Println("[DEBUG] AreaHandler.SearchDetail dipanggil")

	if h.usecase == nil {
		log.Println("[CRITICAL ERROR] h.usecase bernilai NIL! Wire injection gagal.")
		api.Error(c, http.StatusInternalServerError, "Internal Server Error: Usecase nil")
		return
	}

	var req dto.GetAreaDetailRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("[DEBUG] BindQuery Warn/Err: %v", err)
	}

	if req.Search == "" {
		req.Search = c.Query("search")
	}

	details, meta, err := h.usecase.SearchDetail(c.Request.Context(), req)
	if err != nil {
		log.Printf("[ERROR] Usecase SearchDetail Err: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), details, meta)
}
