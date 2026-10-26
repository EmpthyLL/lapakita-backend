package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/wallet/dto"
)

var ErrInvalidPayoutMethod = errors.New("invalid payout method")

func (u *WalletUsecase) CreatePayoutMethod(ctx context.Context, userID uuid.UUID, req dto.CreatePayoutMethodRequest) error {
	method := &entity.PayoutMethod{
		UserID:            userID,
		PayoutType:        req.PayoutType,
		ProviderCode:      strings.TrimSpace(req.ProviderCode),
		ProviderName:      strings.TrimSpace(req.ProviderName),
		AccountNumber:     strings.TrimSpace(req.AccountNumber),
		AccountHolderName: strings.TrimSpace(req.AccountHolderName),
	}
	if method.ProviderCode == "" || method.ProviderName == "" || method.AccountNumber == "" || method.AccountHolderName == "" {
		return ErrInvalidPayoutMethod
	}
	return u.Repo.Create(ctx, method)
}
