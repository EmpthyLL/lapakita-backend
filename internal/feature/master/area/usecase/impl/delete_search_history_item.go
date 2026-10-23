package impl

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/master/area/dto"
	"lapakita-backend/internal/feature/master/area/usecase"
)

func DeleteSearchHistoryItem(ctx context.Context, u *usecase.AreaUsecase, userID string, deviceID string, fullLabel string) error {
	userKey, guestKey := usecase.GetHistoryKeys(userID, deviceID)
	targetKey := userKey
	if targetKey == "" {
		targetKey = guestKey
	}
	if targetKey == "" {
		return nil
	}
	existingItems, _ := u.RDB.LRange(ctx, targetKey, 0, -1).Result()
	for _, oldStr := range existingItems {
		var oldItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(oldStr), &oldItem); err == nil && oldItem.FullLabel == fullLabel {
			u.RDB.LRem(ctx, targetKey, 0, oldStr)
		}
	}
	return nil
}
