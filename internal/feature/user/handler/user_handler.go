package handler

import (
	"net/http"
	"strconv"

	"lapakita-backend/internal/feature/user/dto"
	usermessages "lapakita-backend/internal/feature/user/messages"
	"lapakita-backend/internal/feature/user/usecase"
	"lapakita-backend/internal/feature/user/usecase/impl"
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

	res, err := impl.GetGeneralProfile(c.Request.Context(), h.userUsecase, userID)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserProfileGetSuccess), res)
}

func (h *UserHandler) UpdateGeneralProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.UpdateGeneralProfileRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, err := impl.UpdateGeneralProfile(c.Request.Context(), h.userUsecase, userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserProfileUpdateSuccess), res)
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

	var req dto.GetPhoneNumbersRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, meta, err := impl.GetPhoneNumbers(c.Request.Context(), h.userUsecase, userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, usermessages.KeyUserPhoneGetSuccess), res, meta)
}

func (h *UserHandler) AddPhoneNumber(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.AddPhoneNumberRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := impl.AddPhoneNumber(c.Request.Context(), h.userUsecase, userID, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusCreated, i18n.T(c, usermessages.KeyUserPhoneAddSuccess))
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
		api.Error(c, http.StatusBadRequest, i18n.T(c, usermessages.KeyUserPhoneIndexInvalid))
		return
	}

	var req dto.UpdatePhoneNumberRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := impl.UpdatePhoneNumber(c.Request.Context(), h.userUsecase, userID, index, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserPhoneUpdateSuccess))
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
		api.Error(c, http.StatusBadRequest, i18n.T(c, usermessages.KeyUserPhoneIndexInvalid))
		return
	}

	if err := impl.DeletePhoneNumber(c.Request.Context(), h.userUsecase, userID, index); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserPhoneDeleteSuccess))
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
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := impl.UpdatePassword(c.Request.Context(), h.userUsecase, userID, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserPasswordChangeSuccess))
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
	res, err := impl.GetPersonaProfile(c.Request.Context(), h.userUsecase, userID, role)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserPersonaGetSuccess), res)
}

func (h *UserHandler) UpdatePersonaProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	role := c.Param("role")
	var req dto.UpdatePersonaRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, err := impl.UpdatePersonaProfile(c.Request.Context(), h.userUsecase, userID, role, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserPersonaUpdateSuccess), res)
}

// -----------------------------------------------------------------------------
// 5. LEGAL IDENTITY DOCUMENT
// -----------------------------------------------------------------------------

func (h *UserHandler) GetDocument(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.GetDocumentRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, meta, err := impl.GetDocument(c.Request.Context(), h.userUsecase, userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, usermessages.KeyUserDocumentGetSuccess), res, meta)
}

func (h *UserHandler) UploadDocument(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.UploadDocumentRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := impl.UploadDocument(c.Request.Context(), h.userUsecase, userID, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusCreated, i18n.T(c, usermessages.KeyUserDocumentUploadSuccess))
}

func (h *UserHandler) DeleteDocument(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	docIDStr := c.Param("id")
	docID, err := api.DecodeRequest(docIDStr)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := impl.DeleteDocument(c.Request.Context(), h.userUsecase, userID, docID); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, usermessages.KeyUserDocumentDeleteSuccess))
}
