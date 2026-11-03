package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/internal/feature/business/usecase"
	"lapakita-backend/pkg/api"

	"github.com/gin-gonic/gin"
)

type BusinessHandler struct {
	businessUsecase *usecase.BusinessUsecase
}

func NewBusinessHandler(businessUsecase *usecase.BusinessUsecase) *BusinessHandler {
	return &BusinessHandler{
		businessUsecase: businessUsecase,
	}
}

func (h *BusinessHandler) Create(c *gin.Context) {
	var req dto.CreateBusinessRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid request payload", fields)
		return
	}

	res, err := h.businessUsecase.Create(c.Request.Context(), req)
	if err != nil {
		if err == api.ErrInvalidID {
			api.Error(c, http.StatusBadRequest, "Invalid business type ID format")
			return
		}
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Business created successfully", res)
}

func (h *BusinessHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.Error(c, http.StatusBadRequest, "Business ID is required")
		return
	}

	req := dto.GetBusinessByIDRequest{ID: id}
	lang := api.GetLanguageFromHeader(c)

	res, err := h.businessUsecase.GetByID(c.Request.Context(), req, lang)
	if err != nil {
		if err == api.ErrInvalidID {
			api.Error(c, http.StatusBadRequest, "Invalid business ID format")
			return
		}
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if res == nil {
		api.Error(c, http.StatusNotFound, "Business not found")
		return
	}

	api.Success(c, http.StatusOK, "Business retrieved successfully", res)
}

func (h *BusinessHandler) List(c *gin.Context) {
	var req dto.ListBusinessRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid query parameters", fields)
		return
	}

	lang := api.GetLanguageFromHeader(c)
	items, meta, err := h.businessUsecase.List(c.Request.Context(), req, lang)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, "Businesses retrieved successfully", items, meta)
}

func (h *BusinessHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.Error(c, http.StatusBadRequest, "Business ID is required")
		return
	}

	var req dto.UpdateBusinessRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid request payload", fields)
		return
	}
	req.ID = id

	res, err := h.businessUsecase.Update(c.Request.Context(), req)
	if err != nil {
		if err == api.ErrInvalidID {
			api.Error(c, http.StatusBadRequest, "Invalid ID format")
			return
		}
		if err.Error() == "business not found" {
			api.Error(c, http.StatusNotFound, "Business not found")
			return
		}
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Business updated successfully", res)
}

func (h *BusinessHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		api.Error(c, http.StatusBadRequest, "Business ID is required")
		return
	}

	err := h.businessUsecase.Delete(c.Request.Context(), id)
	if err != nil {
		if err == api.ErrInvalidID {
			api.Error(c, http.StatusBadRequest, "Invalid business ID format")
			return
		}
		if err.Error() == "business not found" {
			api.Error(c, http.StatusNotFound, "Business not found")
			return
		}
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Business deleted successfully", nil)
}
