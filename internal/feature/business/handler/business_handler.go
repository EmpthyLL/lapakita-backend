package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/business/dto"
	"lapakita-backend/internal/feature/business/usecase"
	"lapakita-backend/internal/middleware"
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
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.CreateBusinessRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid request payload", fields)
		return
	}
	req.UserID = userID

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

func (h *BusinessHandler) List(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var req dto.ListBusinessRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid query parameters", fields)
		return
	}
	req.UserID = userID

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

func (h *BusinessHandler) GetFronts(c *gin.Context) {
	businessID := c.Param("id")
	if businessID == "" {
		api.Error(c, http.StatusBadRequest, "Business ID is required")
		return
	}

	var req dto.GetBusinessFrontsRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid query parameters", fields)
		return
	}
	req.BusinessID = businessID

	items, meta, err := h.businessUsecase.GetFronts(c.Request.Context(), req)
	if err != nil {
		switch err.Error() {
		case api.ErrInvalidID.Error():
			api.Error(c, http.StatusBadRequest, "Invalid business ID format")
		case "business not found":
			api.Error(c, http.StatusNotFound, err.Error())
		default:
			api.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, "Business fronts retrieved successfully", items, meta)
}

func (h *BusinessHandler) AddFront(c *gin.Context) {
	businessID := c.Param("id")
	if businessID == "" {
		api.Error(c, http.StatusBadRequest, "Business ID is required")
		return
	}

	var req dto.AddBusinessFrontRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid request payload", fields)
		return
	}
	req.BusinessID = businessID

	res, err := h.businessUsecase.AddFront(c.Request.Context(), req)
	if err != nil {
		switch err.Error() {
		case api.ErrInvalidID.Error():
			api.Error(c, http.StatusBadRequest, "Invalid business ID format")
		case "business not found":
			api.Error(c, http.StatusNotFound, err.Error())
		default:
			api.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	api.Success(c, http.StatusCreated, "Business front added successfully", res)
}

func (h *BusinessHandler) UpdateFront(c *gin.Context) {
	frontID := c.Param("front_id")
	if frontID == "" {
		api.Error(c, http.StatusBadRequest, "Front ID is required")
		return
	}

	var req dto.UpdateBusinessFrontRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		fields := api.BindErrorFields(err)
		api.ErrorWithFields(c, http.StatusBadRequest, "Invalid request payload", fields)
		return
	}
	req.ID = frontID

	res, err := h.businessUsecase.UpdateFront(c.Request.Context(), req)
	if err != nil {
		switch err.Error() {
		case api.ErrInvalidID.Error():
			api.Error(c, http.StatusBadRequest, "Invalid front ID format")
		case "business front not found":
			api.Error(c, http.StatusNotFound, err.Error())
		default:
			api.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	api.Success(c, http.StatusOK, "Business front updated successfully", res)
}

func (h *BusinessHandler) DeleteFront(c *gin.Context) {
	frontID := c.Param("front_id")
	if frontID == "" {
		api.Error(c, http.StatusBadRequest, "Front ID is required")
		return
	}

	err := h.businessUsecase.DeleteFront(c.Request.Context(), frontID)
	if err != nil {
		switch err.Error() {
		case api.ErrInvalidID.Error():
			api.Error(c, http.StatusBadRequest, "Invalid front ID format")
		case "business front not found":
			api.Error(c, http.StatusNotFound, err.Error())
		default:
			api.Error(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	api.Success(c, http.StatusOK, "Business front deleted successfully", nil)
}
