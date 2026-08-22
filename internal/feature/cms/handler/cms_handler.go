package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/cms/dto"
	"lapakita-backend/internal/feature/cms/usecase"
	"lapakita-backend/pkg/api"

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

	var req dto.FAQQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	faqs, err := h.cmsUsecase.GetGroupedFAQs(c.Request.Context(), lang, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FAQs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": faqs})
}

func (h *CMSHandler) GetLegalDocument(c *gin.Context) {
	lang := api.GetLanguageFromHeader(c)

	var req dto.LegalDocumentQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doc, err := h.cmsUsecase.GetLegalDocument(c.Request.Context(), lang, &req)
	if err != nil || doc == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Legal document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": doc})
}
