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
var ErrPayoutMethodAlreadyExists = errors.New("payout method already exists")

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
	exists, err := u.Repo.Exists(ctx, method.UserID, method.PayoutType, method.ProviderCode, method.AccountNumber)
	if err != nil {
		return err
	}
	if exists {
		return ErrPayoutMethodAlreadyExists
	}
	return u.Repo.Create(ctx, method)
}
