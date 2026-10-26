package handler

import (
	"errors"
	"log"
	"net/http"

	"lapakita-backend/internal/feature/wallet/dto"
	"lapakita-backend/internal/feature/wallet/usecase"
	"lapakita-backend/internal/middleware"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/i18n"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WalletHandler struct {
	usecase *usecase.WalletUsecase
}

func NewWalletHandler(usecase *usecase.WalletUsecase) *WalletHandler {
	return &WalletHandler{usecase: usecase}
}

func (h *WalletHandler) GetPaymentMethods(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.GetPayoutMethodsRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyQueryInvalid), api.BindErrorFields(err))
		return
	}

	methods, meta, err := h.usecase.GetPayoutMethods(c.Request.Context(), userID, req)
	if err != nil {
		log.Printf("get payout methods failed: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyInternalServerError))
		return
	}
	api.SuccessWithPagination(c, http.StatusOK, i18n.T(c, i18n.KeyPayoutMethodGetSuccess), methods, meta)
}

func (h *WalletHandler) CreatePaymentMethod(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}

	var req dto.CreatePayoutMethodRequest
	if err := api.BindAndValidate(c, &req); err != nil {
		api.ErrorWithFields(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload), api.BindErrorFields(err))
		return
	}
	if err := h.usecase.CreatePayoutMethod(c.Request.Context(), userID, req); err != nil {
		if errors.Is(err, usecase.ErrInvalidPayoutMethod) {
			api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
			return
		}
		log.Printf("create payout method failed: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyPayoutMethodFailedToCreate))
		return
	}
	api.Success(c, http.StatusCreated, i18n.T(c, i18n.KeyPayoutMethodCreateSuccess), nil)
}

func (h *WalletHandler) DeletePaymentMethod(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}
	methodID, err := api.DecodeRequest(c.Param("id"))
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	err = h.usecase.DeletePayoutMethod(c.Request.Context(), userID, methodID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyPayoutMethodNotFound))
			return
		}
		log.Printf("delete payout method failed: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyPayoutMethodFailedToDelete))
		return
	}
	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyPayoutMethodDeleteSuccess), nil)
}

func (h *WalletHandler) SetPrimaryPaymentMethod(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		api.Error(c, http.StatusUnauthorized, i18n.T(c, i18n.KeyUnauthorized))
		return
	}
	methodID, err := api.DecodeRequest(c.Param("id"))
	if err != nil {
		api.Error(c, http.StatusBadRequest, i18n.T(c, i18n.KeyInvalidPayload))
		return
	}

	err = h.usecase.SetPrimaryPayoutMethod(c.Request.Context(), userID, methodID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			api.Error(c, http.StatusNotFound, i18n.T(c, i18n.KeyPayoutMethodNotFound))
			return
		}
		log.Printf("set primary payout method failed: %v", err)
		api.Error(c, http.StatusInternalServerError, i18n.T(c, i18n.KeyPayoutMethodFailedToSetPrimary))
		return
	}
	api.Success(c, http.StatusOK, i18n.T(c, i18n.KeyPayoutMethodSetPrimarySuccess), nil)
}
