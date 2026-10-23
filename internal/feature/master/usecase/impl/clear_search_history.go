package impl

import (
	"context"
	"lapakita-backend/internal/feature/master/usecase"
)

func ClearSearchHistory(ctx context.Context, u *usecase.MasterUsecase, userID string, deviceID string) error {
	userKey, guestKey := usecase.GetHistoryKeys(userID, deviceID)
	if userKey != "" {
		u.RDB.Del(ctx, userKey)
	}
	if guestKey != "" {
		u.RDB.Del(ctx, guestKey)
	}
	return nil
}
