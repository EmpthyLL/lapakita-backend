package impl

import (
	"context"

	"github.com/google/uuid"
	"lapakita-backend/internal/feature/wallet/usecase"
)

func SetPrimaryPayoutMethod(ctx context.Context, u *usecase.PayoutMethodUsecase, userID, methodID uuid.UUID) error {
	return u.Repo.SetPrimary(ctx, userID, methodID)
}
