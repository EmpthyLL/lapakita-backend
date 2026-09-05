package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/public/dto"
	"lapakita-backend/internal/feature/public/usecase"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
	publicUsecase *usecase.PublicUsecase
}

func NewPublicHandler(publicUsecase *usecase.PublicUsecase) *PublicHandler {
	return &PublicHandler{
		publicUsecase: publicUsecase,
	}
}

func (h *PublicHandler) GetFAQs(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	roleType := c.Param("role_type")

	var req dto.FAQQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid))
		return
	}

	faqs, err := h.publicUsecase.GetGroupedFAQs(c.Request.Context(), lang, roleType, &req)
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

func (h *PublicHandler) GetLegalDocument(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	docType := c.Param("doc_type")

	var req dto.LegalDocumentQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid))
		return
	}

	res, err := h.publicUsecase.GetLegalDocument(c.Request.Context(), lang, docType, &req)
	if err != nil || res == nil {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyPublicLegalNotFound))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyPublicLegalGetSuccess), res)
}

func (h *PublicHandler) SubmitContactInquiry(c *gin.Context) {
	var req dto.SubmitContactInquiryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.publicUsecase.SubmitContactInquiry(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyPublicContactInquirySubmitFailed))
		return
	}

	api.Success(c, http.StatusCreated, i18n.T(c, i18n.KeyPublicContactInquirySubmitSuccess), res)
}
