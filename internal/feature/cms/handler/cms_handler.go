package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/cms/dto"
	"lapakita-backend/internal/feature/cms/usecase"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type CMSHandler struct {
	cmsUsecase *usecase.CMSUsecase
}

func NewCMSHandler(cmsUsecase *usecase.CMSUsecase) *CMSHandler {
	return &CMSHandler{
		cmsUsecase: cmsUsecase,
	}
}

func (h *CMSHandler) GetFAQs(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	roleType := c.Param("role_type")

	var req dto.FAQQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid))
		return
	}

	faqs, err := h.cmsUsecase.GetGroupedFAQs(c.Request.Context(), lang, roleType, &req)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyCMSFAQFailedToGet))
		return
	}

	if len(faqs) == 0 {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyCMSFAQNotFound))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyCMSFAQGetSuccess), faqs)
}

func (h *CMSHandler) GetLegalDocument(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)
	docType := c.Param("doc_type")

	var req dto.LegalDocumentQueryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid))
		return
	}

	res, err := h.cmsUsecase.GetLegalDocument(c.Request.Context(), lang, docType, &req)
	if err != nil || res == nil {
		api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyCMSLegalNotFound))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyCMSLegalGetSuccess), res)
}
