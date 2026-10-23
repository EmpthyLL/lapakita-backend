package impl

import (
	"context"

	"github.com/google/uuid"
	"lapakita-backend/internal/feature/wallet/usecase"
)

func DeletePayoutMethod(ctx context.Context, u *usecase.WalletUsecase, userID, methodID uuid.UUID) error {
	return u.Repo.Delete(ctx, userID, methodID)
}
