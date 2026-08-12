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
	"lapakita-backend/pkg/api"

	"github.com/redis/go-redis/v9"
)

type GeoapifyClient interface {
	SearchAutocomplete(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponseData, api.PaginationMeta, error)
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

func (c *geoapifyClient) SearchAutocomplete(ctx context.Context, req dto.GetAreaRequest, lang string) ([]dto.AreaResponseData, api.PaginationMeta, error) {
	if lang == "" {
		lang = "en"
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit

	// 1. Check Redis Cache
	cacheKey := fmt.Sprintf("geoapify:autocomplete:%s:%s:%d:%d", lang, req.Search, req.Page, req.Limit)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedResult struct {
			Data []dto.AreaResponseData `json:"data"`
			Meta api.PaginationMeta     `json:"meta"`
		}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err == nil {
			return cachedResult.Data, cachedResult.Meta, nil
		}
	}

	// 2. Fetch Geoapify API
	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("invalid url: %w", err)
	}

	fetchLimit := req.Limit + 1

	query := reqURL.Query()
	query.Set("text", req.Search)
	query.Set("limit", fmt.Sprintf("%d", fetchLimit))
	query.Set("offset", fmt.Sprintf("%d", offset))
	query.Set("filter", "countrycode:id")
	query.Set("lang", lang)
	query.Set("apiKey", c.cfg.GeoapifyAPIKey)
	reqURL.RawQuery = query.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("failed to call geoapify service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, api.PaginationMeta{}, fmt.Errorf("geoapify service error (%d): %s", resp.StatusCode, string(body))
	}

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("failed to decode response payload: %w", err)
	}

	// 3. Detect HasNextPage
	hasNextPage := false
	features := geoResp.Features

	if len(features) > req.Limit {
		hasNextPage = true
		features = features[:req.Limit]
	}

	// 4. Manual Object Mapping
	var areas []dto.AreaResponseData
	for _, feature := range features {
		props := feature.Properties

		area := dto.AreaResponseData{
			Formatted:      props.Formatted,
			AddressLine1:   props.AddressLine1,
			AddressLine2:   props.AddressLine2,
			GoogleMapURL:   fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", props.Lat, props.Lon),
			GoogleEmbedURL: fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&hl=%s&z=16&output=embed", props.Lat, props.Lon, lang),
		}

		areas = append(areas, area)
	}

	meta := api.PaginationMeta{
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: hasNextPage,
		HasPrevPage: req.Page > 1,
	}

	// 5. Store Cache
	cachePayload := struct {
		Data []dto.AreaResponseData `json:"data"`
		Meta api.PaginationMeta     `json:"meta"`
	}{
		Data: areas,
		Meta: meta,
	}

	if marshalled, err := json.Marshal(cachePayload); err == nil {
		c.rdb.Set(ctx, cacheKey, string(marshalled), 24*time.Hour)
	}

	return areas, meta, nil
}
