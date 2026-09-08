package handler

import (
	"net/http"
	"strconv"

	"lapakita-backend/internal/feature/user/dto"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// -----------------------------------------------------------------------------
// 1. GENERAL PROFILE
// -----------------------------------------------------------------------------

func (h *UserHandler) GetGeneralProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	res, err := h.userUsecase.GetGeneralProfile(c.Request.Context(), userID)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserProfileGetSuccess), res)
}

func (h *UserHandler) UpdateGeneralProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.UpdateGeneralProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.userUsecase.UpdateGeneralProfile(c.Request.Context(), userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserProfileUpdateSuccess), res)
}

// -----------------------------------------------------------------------------
// 2. PHONE NUMBERS
// -----------------------------------------------------------------------------

func (h *UserHandler) GetPhoneNumbers(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	res, err := h.userUsecase.GetPhoneNumbers(c.Request.Context(), userID)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserPhoneGetSuccess), res)
}

func (h *UserHandler) AddPhoneNumber(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.AddPhoneNumberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.userUsecase.AddPhoneNumber(c.Request.Context(), userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusCreated, i18n.T(c, i18n.KeyUserPhoneAddSuccess), res)
}

func (h *UserHandler) UpdatePhoneNumber(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyUserPhoneIndexInvalid))
		return
	}

	var req dto.UpdatePhoneNumberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.userUsecase.UpdatePhoneNumber(c.Request.Context(), userID, index, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserPhoneUpdateSuccess), res)
}

func (h *UserHandler) DeletePhoneNumber(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	indexStr := c.Param("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyUserPhoneIndexInvalid))
		return
	}

	if err := h.userUsecase.DeletePhoneNumber(c.Request.Context(), userID, index); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserPhoneDeleteSuccess))
}

// -----------------------------------------------------------------------------
// 3. SECURITY / PASSWORD
// -----------------------------------------------------------------------------

func (h *UserHandler) UpdatePassword(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := h.userUsecase.UpdatePassword(c.Request.Context(), userID, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserPasswordChangeSuccess))
}

// -----------------------------------------------------------------------------
// 4. PERSONA PROFILE
// -----------------------------------------------------------------------------

func (h *UserHandler) GetPersonaProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	role := c.Param("role")
	res, err := h.userUsecase.GetPersonaProfile(c.Request.Context(), userID, role)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserPersonaGetSuccess), res)
}

func (h *UserHandler) UpdatePersonaProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	role := c.Param("role")
	var req dto.UpdatePersonaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.userUsecase.UpdatePersonaProfile(c.Request.Context(), userID, role, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserPersonaUpdateSuccess), res)
}

// -----------------------------------------------------------------------------
// 5. LEGAL IDENTITY DOCUMENT
// -----------------------------------------------------------------------------

func (h *UserHandler) UploadDocument(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.UploadDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.userUsecase.UploadDocument(c.Request.Context(), userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusCreated, i18n.T(c, i18n.KeyUserDocumentUploadSuccess), res)
}

func (h *UserHandler) DeleteDocument(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	if err := h.userUsecase.DeleteDocument(c.Request.Context(), userID); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyUserDocumentDeleteSuccess))
}
