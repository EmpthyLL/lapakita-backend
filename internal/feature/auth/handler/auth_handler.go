package handler

import (
	"net/http"

	"lapakita-backend/internal/feature/auth/dto"
	"lapakita-backend/internal/feature/auth/usecase"
	"lapakita-backend/internal/feature/auth/usecase/impl"
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

	if err := impl.Register(c.Request.Context(), h.usecase, req); err != nil {
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

	res, err := impl.Login(c.Request.Context(), h.usecase, req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
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

	authRes, err := impl.GoogleAuth(c.Request.Context(), h.usecase, req)
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

	res, err := impl.CompleteProfile(c.Request.Context(), h.usecase, userID, req)
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

	if err := impl.SendOTP(c.Request.Context(), h.usecase, req); err != nil {
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

	res, err := impl.VerifyOTP(c.Request.Context(), h.usecase, req)
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

	if err := impl.ResetPassword(c.Request.Context(), h.usecase, email, req); err != nil {
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

	res, err := impl.RefreshToken(c.Request.Context(), h.usecase, req)
	if err != nil {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.MessageKey(err.Error())))
		return
	}

	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyTokenRefreshSuccess), res)
}
