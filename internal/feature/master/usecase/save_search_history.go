package usecase

import (
	"context"
	"encoding/json"
	"time"

	"lapakita-backend/internal/feature/master/dto"
)

func (u *MasterUsecase) SaveSearchHistory(ctx context.Context, userID string, deviceID string, req dto.SaveHistoryRequest) error {
	userKey, guestKey := GetHistoryKeys(userID, deviceID)
	targetKey := userKey
	if targetKey == "" {
		targetKey = guestKey
	}
	if targetKey == "" {
		return nil
	}

	historyItem := dto.AreaHistoryItemResponse{
		AreaGeneralResponseData: req.AreaGeneralResponseData,
		SearchedAt:              time.Now().Format(time.RFC3339),
	}
	itemBytes, err := json.Marshal(historyItem)
	if err != nil {
		return err
	}
	itemStr := string(itemBytes)
	existingItems, _ := u.RDB.LRange(ctx, targetKey, 0, -1).Result()
	for _, oldStr := range existingItems {
		var oldItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(oldStr), &oldItem); err == nil && oldItem.FullLabel == req.FullLabel {
			u.RDB.LRem(ctx, targetKey, 0, oldStr)
		}
	}
	if err := u.RDB.LPush(ctx, targetKey, itemStr).Err(); err != nil {
		return err
	}
	if targetKey == userKey {
		u.RDB.Expire(ctx, userKey, UserHistoryTTL)
	} else {
		u.RDB.Expire(ctx, guestKey, GuestHistoryTTL)
	}
	return nil
}
