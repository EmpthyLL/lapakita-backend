package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
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
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := h.usecase.Register(c.Request.Context(), req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success[any](c, http.StatusOK, i18n.T(c, i18n.KeyOTPSendSuccess))
}

// POST /api/v1/auth/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.usecase.Login(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyLoginSuccess), res)
}

// POST /api/v1/auth/google
func (h *AuthHandler) GoogleAuth(c *gin.Context) {
	var req dto.GoogleAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	authRes, setupRes, err := h.usecase.GoogleAuth(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	if setupRes != nil {
		api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyGoogleAuthProfileIncomplete), setupRes)
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyLoginSuccess), authRes)
}

// POST /api/v1/auth/google/complete
func (h *AuthHandler) CompleteProfile(c *gin.Context) {
	var req dto.CompleteProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.usecase.CompleteProfile(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyProfileCompleteSuccess), res)
}

// POST /api/v1/auth/otp/send
func (h *AuthHandler) SendOTP(c *gin.Context) {
	var req dto.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := h.usecase.SendOTP(c.Request.Context(), req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success[any](c, http.StatusOK, i18n.T(c, i18n.KeyOTPSendSuccess))
}

// POST /api/v1/auth/otp/verify
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.usecase.VerifyOTP(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyOTPVerifySuccess), res)
}

// POST /api/v1/auth/reset-password
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	email := c.Query("email")
	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil || email == "" {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	if err := h.usecase.ResetPassword(c.Request.Context(), email, req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success[any](c, http.StatusOK, i18n.T(c, i18n.KeyResetPasswordSuccess))
}

// POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	res, err := h.usecase.RefreshToken(c.Request.Context(), req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyTokenRefreshSuccess), res)
}
