package impl

import (
	"context"
	"encoding/json"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/internal/feature/master/usecase"

	"github.com/redis/go-redis/v9"
)

func GetSearchHistory(ctx context.Context, u *usecase.MasterUsecase, userID string, deviceID string) ([]dto.AreaHistoryItemResponse, error) {
	userKey, guestKey := usecase.GetHistoryKeys(userID, deviceID)
	if userKey != "" && guestKey != "" {
		usecase.MergeGuestToUserHistory(ctx, u.RDB, userKey, guestKey)
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
		u.RDB.Expire(ctx, userKey, usecase.UserHistoryTTL)
	} else if guestKey != "" {
		u.RDB.Expire(ctx, guestKey, usecase.GuestHistoryTTL)
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
