package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"lapakita-backend/config"
	"lapakita-backend/internal/feature/area/dto"
	"lapakita-backend/pkg/api"

	"github.com/redis/go-redis/v9"
)

type GeoapifyClient interface {
	SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error)
	SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error)
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

func cleanComponent(val string) string {
	cleaned := strings.TrimSpace(val)
	cleaned = strings.TrimPrefix(cleaned, "Kota ")
	cleaned = strings.TrimPrefix(cleaned, "Kabupaten ")
	cleaned = strings.TrimPrefix(cleaned, "City of ")
	cleaned = strings.TrimPrefix(cleaned, "Provinsi ")
	return cleaned
}

// -----------------------------------------------------------------------------
// 1. SEARCH GENERAL (Pencarian General Dropdown UI)
// -----------------------------------------------------------------------------
func (c *geoapifyClient) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit

	cacheKey := fmt.Sprintf("geoapify:general:%s:%d:%d", req.Search, req.Page, req.Limit)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedResult struct {
			Data []dto.AreaGeneralResponseData `json:"data"`
			Meta api.PaginationMeta            `json:"meta"`
		}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err == nil {
			return cachedResult.Data, cachedResult.Meta, nil
		}
	}

	fetchLimit := req.Limit + 1
	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("invalid url: %w", err)
	}

	query := reqURL.Query()
	query.Set("text", req.Search)
	query.Set("limit", fmt.Sprintf("%d", fetchLimit))
	query.Set("offset", fmt.Sprintf("%d", offset))
	query.Set("filter", "countrycode:id")
	query.Set("lang", "id")
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
		return nil, api.PaginationMeta{}, fmt.Errorf("geoapify error (%d): %s", resp.StatusCode, string(body))
	}

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("failed to decode response payload: %w", err)
	}

	hasNextPage := false
	features := geoResp.Features

	if len(features) > req.Limit {
		hasNextPage = true
		features = features[:req.Limit]
	}

	var items []dto.AreaGeneralResponseData
	for _, f := range features {
		p := f.Properties

		cityClean := cleanComponent(p.City)
		stateClean := cleanComponent(p.State)
		country := p.Country
		if country == "" {
			country = "Indonesia"
		}

		entityType := "street"
		title := p.Street
		if title == "" {
			title = p.Name
		}

		var subtitleParts []string

		switch p.ResultType {
		case "country":
			entityType = "country"
			title = country
			subtitleParts = []string{"Negara"}
		case "state", "province":
			entityType = "province"
			title = stateClean
			subtitleParts = []string{country}
		case "city", "county":
			entityType = "city"
			title = cityClean
			if stateClean != "" {
				subtitleParts = append(subtitleParts, stateClean)
			}
			subtitleParts = append(subtitleParts, country)
		case "suburb", "district":
			entityType = "district"
			if p.Suburb != "" {
				title = p.Suburb
			} else {
				title = p.District
			}
			if cityClean != "" {
				subtitleParts = append(subtitleParts, cityClean)
			}
			if stateClean != "" {
				subtitleParts = append(subtitleParts, stateClean)
			}
			subtitleParts = append(subtitleParts, country)
		default:
			entityType = "street"
			if title == "" {
				title = p.Formatted
			}
			if p.Suburb != "" {
				subtitleParts = append(subtitleParts, p.Suburb)
			}
			if cityClean != "" {
				subtitleParts = append(subtitleParts, cityClean)
			}
			if stateClean != "" {
				subtitleParts = append(subtitleParts, stateClean)
			}
			subtitleParts = append(subtitleParts, country)
		}

		subtitle := strings.Join(subtitleParts, ", ")
		fullLabel := title
		if subtitle != "" && entityType != "country" {
			fullLabel = fmt.Sprintf("%s, %s", title, subtitle)
		}

		item := dto.AreaGeneralResponseData{
			Type:        entityType,
			Title:       title,
			Subtitle:    subtitle,
			FullLabel:   fullLabel,
			Country:     country,
			CountryCode: strings.ToUpper(p.CountryCode),
			City:        cityClean,
			Province:    stateClean,
		}

		items = append(items, item)
	}

	meta := api.PaginationMeta{
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: hasNextPage,
		HasPrevPage: req.Page > 1,
	}

	cachePayload := struct {
		Data []dto.AreaGeneralResponseData `json:"data"`
		Meta api.PaginationMeta            `json:"meta"`
	}{
		Data: items,
		Meta: meta,
	}

	if marshalled, err := json.Marshal(cachePayload); err == nil {
		c.rdb.Set(ctx, cacheKey, string(marshalled), 24*time.Hour)
	}

	return items, meta, nil
}

// -----------------------------------------------------------------------------
// 2. SEARCH DETAIL (Untuk Auto-fill Form Input Owner)
// -----------------------------------------------------------------------------
func (c *geoapifyClient) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit

	cacheKey := fmt.Sprintf("geoapify:detail:%s:%d:%d", req.Search, req.Page, req.Limit)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedResult struct {
			Data []dto.AreaDetailResponseData `json:"data"`
			Meta api.PaginationMeta           `json:"meta"`
		}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err == nil {
			return cachedResult.Data, cachedResult.Meta, nil
		}
	}

	fetchLimit := req.Limit + 1
	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("invalid url: %w", err)
	}

	query := reqURL.Query()
	query.Set("text", req.Search)
	query.Set("limit", fmt.Sprintf("%d", fetchLimit))
	query.Set("offset", fmt.Sprintf("%d", offset))
	query.Set("filter", "countrycode:id")
	query.Set("lang", "id")
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
		return nil, api.PaginationMeta{}, fmt.Errorf("geoapify error (%d): %s", resp.StatusCode, string(body))
	}

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return nil, api.PaginationMeta{}, fmt.Errorf("failed to decode response payload: %w", err)
	}

	hasNextPage := false
	features := geoResp.Features

	if len(features) > req.Limit {
		hasNextPage = true
		features = features[:req.Limit]
	}

	var items []dto.AreaDetailResponseData
	for _, f := range features {
		p := f.Properties

		street := p.Street
		if street == "" {
			street = p.Name
		}

		country := p.Country
		if country == "" {
			country = "Indonesia"
		}

		item := dto.AreaDetailResponseData{
			Formatted:      p.Formatted,
			StreetAddress:  street,
			Suburb:         p.Suburb,
			District:       p.District,
			City:           cleanComponent(p.City),
			Province:       cleanComponent(p.State),
			Country:        country,
			CountryCode:    strings.ToUpper(p.CountryCode),
			PostalCode:     p.Postcode,
			Latitude:       p.Lat,
			Longitude:      p.Lon,
			MapURL:         fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", p.Lat, p.Lon),
			EmbeddedMapURL: fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&hl=id&z=16&output=embed", p.Lat, p.Lon),
		}

		items = append(items, item)
	}

	meta := api.PaginationMeta{
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: hasNextPage,
		HasPrevPage: req.Page > 1,
	}

	cachePayload := struct {
		Data []dto.AreaDetailResponseData `json:"data"`
		Meta api.PaginationMeta           `json:"meta"`
	}{
		Data: items,
		Meta: meta,
	}

	if marshalled, err := json.Marshal(cachePayload); err == nil {
		c.rdb.Set(ctx, cacheKey, string(marshalled), 24*time.Hour)
	}

	return items, meta, nil
}
