package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"lapakita-backend/config"
	"lapakita-backend/internal/feature/area/dto"

	"github.com/redis/go-redis/v9"
)

type GeoapifyRawResponse struct {
	Features []struct {
		Properties dto.AreaResponse `json:"properties"`
	} `json:"features"`
}

type GeoapifyClient interface {
	FetchPlaces(ctx context.Context, req dto.GetAreaRequest) ([]dto.AreaResponse, error)
}

type geoapifyClient struct {
	cfg        *config.Config
	rdb        *redis.Client
	httpClient *http.Client
}

func NewGeoapifyClient(cfg *config.Config, rdb *redis.Client) GeoapifyClient {
	return &geoapifyClient{
		cfg:        cfg,
		rdb:        rdb,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *geoapifyClient) FetchPlaces(ctx context.Context, req dto.GetAreaRequest) ([]dto.AreaResponse, error) {
	// 1. Check Redis Cache
	cacheKey := fmt.Sprintf("geoapify:places:%s:%.4f:%.4f:%d", req.Categories, req.Lat, req.Lon, req.Radius)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedPlaces []dto.AreaResponse
		if err := json.Unmarshal([]byte(cachedData), &cachedPlaces); err == nil {
			return cachedPlaces, nil // Cache Hit
		}
	}

	// 2. Cache Miss: Query External Geoapify API
	reqURL, err := url.Parse(c.cfg.GeoapifyBaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	query := reqURL.Query()
	query.Set("categories", req.Categories)
	query.Set("filter", fmt.Sprintf("circle:%f,%f,%d", req.Lon, req.Lat, req.Radius))
	query.Set("bias", fmt.Sprintf("proximity:%f,%f", req.Lon, req.Lat))
	query.Set("limit", "20")
	query.Set("apiKey", c.cfg.GeoapifyAPIKey)
	reqURL.RawQuery = query.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat http request: %w", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("gagal menghubungi Geoapify: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("geoapify error status %d: %s", resp.StatusCode, string(body))
	}

	var geoResp GeoapifyRawResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return nil, fmt.Errorf("gagal decode response json: %w", err)
	}

	var places []dto.AreaResponse
	for _, feature := range geoResp.Features {
		item := feature.Properties

		item.GoogleMapURL = fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", item.Lat, item.Lon)
		item.GoogleEmbedURL = fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&hl=id&z=16&output=embed", item.Lat, item.Lon)

		places = append(places, item)
	}

	// 3. Save to Redis with 24 Hours TTL
	if marshalled, err := json.Marshal(places); err == nil {
		c.rdb.Set(ctx, cacheKey, string(marshalled), 24*time.Hour)
	}

	return places, nil
}
