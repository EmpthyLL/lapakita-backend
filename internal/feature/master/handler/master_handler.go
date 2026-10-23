package handler

import (
	"log"
	"net/http"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/internal/feature/master/usecase"
	"lapakita-backend/internal/feature/master/usecase/impl"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type MasterHandler struct {
	usecase *usecase.MasterUsecase
}

func NewMasterHandler(masterUsecase *usecase.MasterUsecase) *MasterHandler {
	return &MasterHandler{usecase: masterUsecase}
}

func (h *MasterHandler) SearchGeneral(c *gin.Context) {
	if h.usecase == nil {
		api.Error(c, http.StatusInternalServerError, "Internal Server Error: Usecase nil")
		return
	}
	var req dto.GetAreaGeneralRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		log.Printf("bind area search query failed: %v", err)
	}
	if req.Search == "" {
		req.Search = c.Query("search")
	}

	areas, meta, err := impl.SearchGeneral(c.Request.Context(), h.usecase, req)
	if err != nil {
		log.Printf("area search failed: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}
	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), areas, meta)
}

func (h *MasterHandler) SearchDetail(c *gin.Context) {
	if h.usecase == nil {
		api.Error(c, http.StatusInternalServerError, "Internal Server Error: Usecase nil")
		return
	}
	var req dto.GetAreaDetailRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		log.Printf("bind area detail query failed: %v", err)
	}
	if req.Search == "" {
		req.Search = c.Query("search")
	}

	details, meta, err := impl.SearchDetail(c.Request.Context(), h.usecase, req)
	if err != nil {
		log.Printf("area detail search failed: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyAreaSearchFailed))
		return
	}
	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyAreaSearchSuccess), details, meta)
}

func extractAuthAndDevice(c *gin.Context) (userID string, deviceID string) {
	if parsedUUID, ok := middleware.GetUserID(c); ok {
		userID = parsedUUID.String()
	}
	return userID, c.GetHeader("X-Device-ID")
}

func (h *MasterHandler) GetHistory(c *gin.Context) {
	userID, deviceID := extractAuthAndDevice(c)
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

func (h *MasterHandler) SaveHistory(c *gin.Context) {
	userID, deviceID := extractAuthAndDevice(c)
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

func (h *MasterHandler) ClearHistory(c *gin.Context) {
	userID, deviceID := extractAuthAndDevice(c)
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

func (h *MasterHandler) DeleteItemHistory(c *gin.Context) {
	userID, deviceID := extractAuthAndDevice(c)
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

func (h *MasterHandler) GetFAQs(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	roleType := c.Param("role_type")
	var req dto.FAQQueryRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid), api.BindErrorFields(err))
		return
	}
	faqs, err := impl.GetGroupedFAQs(c.Request.Context(), h.usecase, lang, roleType, &req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyPublicFAQFailedToGet))
		return
	}
	if len(faqs) == 0 {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyPublicFAQNotFound))
		return
	}
	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyPublicFAQGetSuccess), faqs)
}

func (h *MasterHandler) GetLegalDocument(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	docType := c.Param("doc_type")
	var req dto.LegalDocumentQueryRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid), api.BindErrorFields(err))
		return
	}
	res, err := impl.GetLegalDocument(c.Request.Context(), h.usecase, lang, docType, &req)
	if err != nil || res == nil {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyPublicLegalNotFound))
		return
	}
	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyPublicLegalGetSuccess), res)
}

func (h *MasterHandler) SubmitContact(c *gin.Context) {
	var req dto.SubmitContactRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}
	res, err := impl.SubmitContact(c.Request.Context(), h.usecase, req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyPublicContactInquirySubmitFailed))
		return
	}
	api.Success(c, http.StatusCreated, i18n.T(c, i18n.KeyPublicContactInquirySubmitSuccess), res)
}

func (h *MasterHandler) GetBusinessTypes(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	var req dto.GetBusinessTypesRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid), api.BindErrorFields(err))
		return
	}
	data, meta, err := impl.GetBusinessTypes(c.Request.Context(), h.usecase, lang, &req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyBusinessTypeFailedToGet))
		return
	}
	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyBusinessTypeGetSuccess), data, meta)
}
