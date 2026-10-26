package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (u *WalletUsecase) SetPrimaryPayoutMethod(ctx context.Context, userID, methodID uuid.UUID) error {
	return u.Repo.SetPrimary(ctx, userID, methodID)
}
