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

type GeoapifyClient interface {
	SearchAutocomplete(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponse, error)
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

func (c *geoapifyClient) SearchAutocomplete(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponse, error) {
	if lang == "" {
		lang = "en"
	}

	// 1. Check Redis Cache
	cacheKey := fmt.Sprintf("geoapify:autocomplete:%s:%s", lang, req.Text)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedAreas []dto.AreaResponse
		if err := json.Unmarshal([]byte(cachedData), &cachedAreas); err == nil {
			return cachedAreas, nil
		}
	}

	// 2. Query External API Geoapify
	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}

	query := reqURL.Query()
	query.Set("text", req.Text)
	query.Set("limit", "10")
	query.Set("filter", "countrycode:id")
	query.Set("lang", lang)
	query.Set("apiKey", c.cfg.GeoapifyAPIKey)
	reqURL.RawQuery = query.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat request: %w", err)
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

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return nil, fmt.Errorf("gagal decode response json: %w", err)
	}

	// 3. Transformasi dari Geoapify ke DTO FE (Hanya ambil yang perlu)
	var areas []dto.AreaResponse
	for _, feature := range geoResp.Features {
		props := feature.Properties

		area := dto.AreaResponse{
			Formatted:      props.Formatted,
			AddressLine1:   props.AddressLine1,
			AddressLine2:   props.AddressLine2,
			GoogleMapURL:   fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", props.Lat, props.Lon),
			GoogleEmbedURL: fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&hl=%s&z=16&output=embed", props.Lat, props.Lon, lang),
		}

		areas = append(areas, area)
	}

	// 4. Save to Cache (24 Hours TTL)
	if marshalled, err := json.Marshal(areas); err == nil {
		c.rdb.Set(ctx, cacheKey, string(marshalled), 24*time.Hour)
	}

	return areas, nil
}
