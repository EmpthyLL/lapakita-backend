package usecase

import (
	"context"
	"lapakita-backend/internal/feature/area/usecase"
)

func (u *AreaUsecase) ClearSearchHistory(ctx context.Context, userID string, deviceID string) error {
	userKey, guestKey := usecase.GetHistoryKeys(userID, deviceID)
	if userKey != "" {
		u.rdb.Del(ctx, userKey)
	}
	if guestKey != "" {
		u.rdb.Del(ctx, guestKey)
	}
	return nil
}
