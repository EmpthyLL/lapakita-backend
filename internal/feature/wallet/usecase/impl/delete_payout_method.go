package impl

import (
	"context"

	"github.com/google/uuid"
	"lapakita-backend/internal/feature/wallet/usecase"
)

func DeletePayoutMethod(ctx context.Context, u *usecase.PayoutMethodUsecase, userID, methodID uuid.UUID) error {
	return u.Repo.Delete(ctx, userID, methodID)
}
