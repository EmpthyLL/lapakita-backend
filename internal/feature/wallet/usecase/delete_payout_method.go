package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (u *WalletUsecase) DeletePayoutMethod(ctx context.Context, userID, methodID uuid.UUID) error {
	return u.Repo.Delete(ctx, userID, methodID)
}
