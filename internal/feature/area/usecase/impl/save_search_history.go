package usecase

import (
	"context"
	"encoding/json"
	"time"

	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/internal/feature/area/usecase"
)

func (u *AreaUsecase) SaveSearchHistory(ctx context.Context, userID string, deviceID string, req dto.SaveHistoryRequest) error {
	userKey, guestKey := usecase.GetHistoryKeys(userID, deviceID)
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
	existingItems, _ := u.rdb.LRange(ctx, targetKey, 0, -1).Result()
	for _, oldStr := range existingItems {
		var oldItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(oldStr), &oldItem); err == nil && oldItem.FullLabel == req.FullLabel {
			u.rdb.LRem(ctx, targetKey, 0, oldStr)
		}
	}
	if err := u.rdb.LPush(ctx, targetKey, itemStr).Err(); err != nil {
		return err
	}
	if targetKey == userKey {
		u.rdb.Expire(ctx, userKey, usecase.UserHistoryTTL)
	} else {
		u.rdb.Expire(ctx, guestKey, usecase.GuestHistoryTTL)
	}
	return nil
}
