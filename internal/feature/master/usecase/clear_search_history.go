package usecase

import (
	"context"
)

func (u *MasterUsecase) ClearSearchHistory(ctx context.Context, userID string, deviceID string) error {
	userKey, guestKey := GetHistoryKeys(userID, deviceID)
	if userKey != "" {
		u.RDB.Del(ctx, userKey)
	}
	if guestKey != "" {
		u.RDB.Del(ctx, guestKey)
	}
	return nil
}
