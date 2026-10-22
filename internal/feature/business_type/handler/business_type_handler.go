package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/business_type/dto"
	"lapakita-backend/internal/feature/business_type/usecase"
	"lapakita-backend/internal/feature/business_type/usecase/impl"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type BusinessTypeHandler struct {
	usecase *usecase.BusinessTypeUsecase
}

func NewBusinessTypeHandler(usecase *usecase.BusinessTypeUsecase) *BusinessTypeHandler {
	return &BusinessTypeHandler{usecase: usecase}
}

func (h *BusinessTypeHandler) GetBusinessTypes(c *gin.Context) {
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
