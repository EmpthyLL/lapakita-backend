package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"lapakita-backend/internal/feature/area/client"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"

	"github.com/redis/go-redis/v9"
)

const (
	GuestHistoryTTL = 30 * 24 * time.Hour
	UserHistoryTTL  = 180 * 24 * time.Hour
)

type AreaUsecase struct {
	geoClient *client.GeoapifyClient
	rdb       *redis.Client
}

func NewAreaUsecase(geoClient *client.GeoapifyClient, rdb *redis.Client) *AreaUsecase {
	return &AreaUsecase{
		geoClient: geoClient,
		rdb:       rdb,
	}
}

func (u *AreaUsecase) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchGeneral(ctx, req)
}

func (u *AreaUsecase) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	return u.geoClient.SearchDetail(ctx, req)
}

func (u *AreaUsecase) helperGetHistoryKeys(userID string, deviceID string) (userKey string, guestKey string) {
	if userID != "" {
		userKey = fmt.Sprintf("search_history:user:%s", userID)
	}
	if deviceID != "" {
		guestKey = fmt.Sprintf("search_history:guest:%s", deviceID)
	}
	return userKey, guestKey
}

// GetSearchHistory mengambil seluruh riwayat pencarian beserta searched_at
func (u *AreaUsecase) GetSearchHistory(ctx context.Context, userID string, deviceID string) ([]dto.AreaHistoryItemResponse, error) {
	userKey, guestKey := u.helperGetHistoryKeys(userID, deviceID)

	if userKey != "" && guestKey != "" {
		u.mergeGuestToUserHistory(ctx, userKey, guestKey)
	}

	targetKey := userKey
	if targetKey == "" {
		targetKey = guestKey
	}

	if targetKey == "" {
		return []dto.AreaHistoryItemResponse{}, nil
	}

	rawItems, err := u.rdb.LRange(ctx, targetKey, 0, -1).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if userKey != "" {
		u.rdb.Expire(ctx, userKey, UserHistoryTTL)
	} else if guestKey != "" {
		u.rdb.Expire(ctx, guestKey, GuestHistoryTTL)
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

// SaveSearchHistory menyimpan 1 item area ke Redis dengan timestamp waktu pencarian
func (u *AreaUsecase) SaveSearchHistory(ctx context.Context, userID string, deviceID string, req dto.SaveHistoryRequest) error {
	userKey, guestKey := u.helperGetHistoryKeys(userID, deviceID)

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

	// 1. Hapus item lama yang punya FullLabel sama agar posisinya terisi yang baru di index 0
	existingItems, _ := u.rdb.LRange(ctx, targetKey, 0, -1).Result()
	for _, oldStr := range existingItems {
		var oldItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(oldStr), &oldItem); err == nil {
			if oldItem.FullLabel == req.FullLabel {
				u.rdb.LRem(ctx, targetKey, 0, oldStr)
			}
		}
	}

	// 2. Push ke urutan paling atas (LPUSH)
	if err := u.rdb.LPush(ctx, targetKey, itemStr).Err(); err != nil {
		return err
	}

	// 3. Reset TTL
	if targetKey == userKey {
		u.rdb.Expire(ctx, userKey, UserHistoryTTL)
	} else {
		u.rdb.Expire(ctx, guestKey, GuestHistoryTTL)
	}

	return nil
}

// DeleteSearchHistoryItem menghapus 1 item spesifik berdasarkan FullLabel
func (u *AreaUsecase) DeleteSearchHistoryItem(ctx context.Context, userID string, deviceID string, fullLabel string) error {
	userKey, guestKey := u.helperGetHistoryKeys(userID, deviceID)

	targetKey := userKey
	if targetKey == "" {
		targetKey = guestKey
	}

	if targetKey == "" {
		return nil
	}

	existingItems, _ := u.rdb.LRange(ctx, targetKey, 0, -1).Result()
	for _, oldStr := range existingItems {
		var oldItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(oldStr), &oldItem); err == nil {
			if oldItem.FullLabel == fullLabel {
				u.rdb.LRem(ctx, targetKey, 0, oldStr)
			}
		}
	}

	return nil
}

// ClearSearchHistory menghapus SELURUH riwayat pencarian
func (u *AreaUsecase) ClearSearchHistory(ctx context.Context, userID string, deviceID string) error {
	userKey, guestKey := u.helperGetHistoryKeys(userID, deviceID)

	if userKey != "" {
		u.rdb.Del(ctx, userKey)
	}
	if guestKey != "" {
		u.rdb.Del(ctx, guestKey)
	}

	return nil
}

func (u *AreaUsecase) mergeGuestToUserHistory(ctx context.Context, userKey string, guestKey string) {
	guestItems, err := u.rdb.LRange(ctx, guestKey, 0, -1).Result()
	if err != nil || len(guestItems) == 0 {
		return
	}

	userItems, _ := u.rdb.LRange(ctx, userKey, 0, -1).Result()

	existMap := make(map[string]bool)
	for _, uStr := range userItems {
		var uItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(uStr), &uItem); err == nil {
			existMap[uItem.FullLabel] = true
		}
	}

	for i := len(guestItems) - 1; i >= 0; i-- {
		gStr := guestItems[i]
		var gItem dto.AreaHistoryItemResponse
		if err := json.Unmarshal([]byte(gStr), &gItem); err == nil {
			if !existMap[gItem.FullLabel] {
				u.rdb.LPush(ctx, userKey, gStr)
				existMap[gItem.FullLabel] = true
			}
		}
	}

	u.rdb.Expire(ctx, userKey, UserHistoryTTL)
	u.rdb.Del(ctx, guestKey)
}
