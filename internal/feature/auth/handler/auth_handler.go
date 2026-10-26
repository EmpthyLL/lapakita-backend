package handler

import (
	"log"
	"net/http"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

// POST /api/v1/auth/register
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := h.usecase.Register(c.Request.Context(), req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyOTPSendSuccess))
}

// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, err := h.usecase.Login(c.Request.Context(), req)
	if err != nil {
		messageKey := i18n.MessageKey(err.Error())
		switch messageKey {
		case i18n.KeyUserNotFound, i18n.KeyWrongPassword:
			api.Error(c, http.StatusUnauthorized, i18n.T(c, messageKey))
		default:
			log.Printf("auth login failed: %v", err)
			api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyInternalServerError))
		}
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyLoginSuccess), res)
}

// POST /api/v1/auth/google
func (h *AuthHandler) GoogleAuth(c *gin.Context) {
	var req dto.GoogleAuthRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	authRes, err := h.usecase.GoogleAuth(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyLoginSuccess), authRes)
}

// PUT /api/v1/auth/complete-profile (Protected Route via JWT Middleware)
func (h *AuthHandler) CompleteProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.CompleteProfileRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, err := h.usecase.CompleteProfile(c.Request.Context(), userID, req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyProfileCompleteSuccess), res)
}

// POST /api/v1/auth/otp/send
func (h *AuthHandler) SendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	if err := h.usecase.SendOTP(c.Request.Context(), req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyOTPSendSuccess))
}

// POST /api/v1/auth/otp/verify
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, err := h.usecase.VerifyOTP(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyOTPVerifySuccess), res)
}

// POST /api/v1/auth/reset-password/:email
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	email := c.Param("email")
	var req dto.ResetPasswordRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}
	if email == "" {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := h.usecase.ResetPassword(c.Request.Context(), email, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyResetPasswordSuccess))
}

// POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}

	res, err := h.usecase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyTokenRefreshSuccess), res)
}
