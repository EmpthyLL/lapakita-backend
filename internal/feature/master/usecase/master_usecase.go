package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"lapakita-backend/internal/feature/master/client"
	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/internal/feature/master/repository"
	"lapakita-backend/pkg/mailer"

	"github.com/redis/go-redis/v9"
)

type MasterUsecase struct {
	Repo      *repository.MasterRepository
	GeoClient *client.GeoapifyClient
	RDB       *redis.Client
	Mailer    *mailer.Mailer
}

func NewMasterUsecase(repo *repository.MasterRepository, geoClient *client.GeoapifyClient, rdb *redis.Client, mailer *mailer.Mailer) *MasterUsecase {
	return &MasterUsecase{
		Repo:      repo,
		GeoClient: geoClient,
		RDB:       rdb,
		Mailer:    mailer,
	}
}

const (
	GuestHistoryTTL = 30 * 24 * time.Hour
	UserHistoryTTL  = 180 * 24 * time.Hour
)

func GetHistoryKeys(userID string, deviceID string) (userKey string, guestKey string) {
	if userID != "" {
		userKey = fmt.Sprintf("search_history:user:%s", userID)
	}
	if deviceID != "" {
		guestKey = fmt.Sprintf("search_history:guest:%s", deviceID)
	}
	return userKey, guestKey
}

func MergeGuestToUserHistory(ctx context.Context, rdb *redis.Client, userKey string, guestKey string) {
	guestItems, err := rdb.LRange(ctx, guestKey, 0, -1).Result()
	if err != nil || len(guestItems) == 0 {
		return
	}
	userItems, _ := rdb.LRange(ctx, userKey, 0, -1).Result()
	existMap := make(map[string]bool)
	for _, value := range userItems {
		var item dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(value), &item); err == nil {
			existMap[item.FullLabel] = true
		}
	}
	for i := len(guestItems) - 1; i >= 0; i-- {
		value := guestItems[i]
		var item dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(value), &item); err == nil && !existMap[item.FullLabel] {
			rdb.LPush(ctx, userKey, value)
			existMap[item.FullLabel] = true
		}
	}
	rdb.Expire(ctx, userKey, UserHistoryTTL)
	rdb.Del(ctx, guestKey)
}

func ResolveLanguage(requested, fallback string) string {
	if requested != "" {
		return requested
	}
	if fallback != "" {
		return fallback
	}
	return "en"
}
