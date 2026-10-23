package handler

import (
	"log"
	"net/http"

	"lapakita-backend/internal/feature/master/area/dto"
	"lapakita-backend/internal/feature/master/area/usecase"
	"lapakita-backend/internal/feature/master/area/usecase/impl"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type AreaHandler struct {
	usecase *usecase.AreaUsecase
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
	if err := api.BindAndValidate(c, &req); err != nil {
		log.Printf("[DEBUG] BindQuery Warn/Err: %v", err)
	}

	// Fallback pembacaan manual jika binding Gin gagal membaca query "search"
	if req.Search == "" {
		req.Search = c.Query("search")
	}

	log.Printf("[DEBUG] Search Query Extracted: '%s'", req.Search)

	areas, meta, err := impl.SearchGeneral(c.Request.Context(), h.usecase, req)
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

	if err := api.BindAndValidate(c, &req); err != nil {
		log.Printf("[DEBUG] BindQuery Warn/Err: %v", err)
	}

	if req.Search == "" {
		req.Search = c.Query("search")
	}

	details, meta, err := impl.SearchDetail(c.Request.Context(), h.usecase, req)
	if err != nil {
		log.Printf("[ERROR] Usecase SearchDetail Err: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), details, meta)
}

func helperExtractAuthAndDevice(c *gin.Context) (userID string, deviceID string) {
	if parsedUUID, ok := middleware.GetUserID(c); ok {
		userID = parsedUUID.String()
	}
	deviceID = c.GetHeader("X-Device-ID")
	return userID, deviceID
}

// GET /api/v1/areas/history
func (h *AreaHandler) GetHistory(c *gin.Context) {
	userID, deviceID := helperExtractAuthAndDevice(c)

	if userID == "" && deviceID == "" {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyAreaDeviceIDRequired))
		return
	}

	history, err := impl.GetSearchHistory(c.Request.Context(), h.usecase, userID, deviceID)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyInternalServerError))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyAreaHistoryGetSuccess), history)
}

// POST /api/v1/areas/history
func (h *AreaHandler) SaveHistory(c *gin.Context) {
	userID, deviceID := helperExtractAuthAndDevice(c)

	if userID == "" && deviceID == "" {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyAreaDeviceIDRequired))
		return
	}

	var req dto.SaveHistoryRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := impl.SaveSearchHistory(c.Request.Context(), h.usecase, userID, deviceID, req); err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyInternalServerError))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyAreaHistorySaveSuccess))
}

// DELETE /api/v1/areas/history
func (h *AreaHandler) ClearHistory(c *gin.Context) {
	userID, deviceID := helperExtractAuthAndDevice(c)

	if userID == "" && deviceID == "" {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyAreaDeviceIDRequired))
		return
	}

	if err := impl.ClearSearchHistory(c.Request.Context(), h.usecase, userID, deviceID); err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyInternalServerError))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyAreaHistoryDeleteSuccess))
}

// DELETE /api/v1/areas/history/item
func (h *AreaHandler) DeleteItemHistory(c *gin.Context) {
	userID, deviceID := helperExtractAuthAndDevice(c)

	if userID == "" && deviceID == "" {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyAreaDeviceIDRequired))
		return
	}

	var req dto.DeleteHistoryItemRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := impl.DeleteSearchHistoryItem(c.Request.Context(), h.usecase, userID, deviceID, req.FullLabel); err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyInternalServerError))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyAreaHistoryDeleteItemSuccess))
}
