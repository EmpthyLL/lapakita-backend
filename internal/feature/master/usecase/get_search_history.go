package usecase

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/master/dto"

	"github.com/redis/go-redis/v9"
)

func (u *MasterUsecase) GetSearchHistory(ctx context.Context, userID string, deviceID string) ([]dto.AreaHistoryItemResponse, error) {
	userKey, guestKey := GetHistoryKeys(userID, deviceID)
	if userKey != "" && guestKey != "" {
		MergeGuestToUserHistory(ctx, u.RDB, userKey, guestKey)
	}
	targetKey := userKey
	if targetKey == "" {
		targetKey = guestKey
	}
	if targetKey == "" {
		return []dto.AreaHistoryItemResponse{}, nil
	}

	rawItems, err := u.RDB.LRange(ctx, targetKey, 0, -1).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if userKey != "" {
		u.RDB.Expire(ctx, userKey, UserHistoryTTL)
	} else if guestKey != "" {
		u.RDB.Expire(ctx, guestKey, GuestHistoryTTL)
	}

	result := make([]dto.AreaHistoryItemResponse, 0, len(rawItems))
	for _, itemStr := range rawItems {
		var item dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(itemStr), &item); err == nil {
			result = append(result, item)
		}
	}
	return result, nil
}
