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

type GeoapifyClient struct {
	cfg        *config.Config
	rdb        *redis.Client
	httpClient *http.Client
}

func NewGeoapifyClient(cfg *config.Config, rdb *redis.Client) *GeoapifyClient {
	return &GeoapifyClient{
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
// 1. SEARCH GENERAL
// -----------------------------------------------------------------------------
func (c *GeoapifyClient) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	cleanSearch := strings.TrimSpace(req.Search)

	// Pastikan mengembalikan slice kosong [] dan BUKAN nil (mencegah response 'null' di JSON FE)
	items := make([]dto.AreaGeneralResponseData, 0)

	if len(cleanSearch) < 2 {
		return items, api.PaginationMeta{
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: false,
		}, nil
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit

	// Cache key aman
	cacheKey := fmt.Sprintf("geoapify:general_v3:%s:%d:%d", cleanSearch, req.Page, req.Limit)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedResult struct {
			Data []dto.AreaGeneralResponseData `json:"data"`
			Meta api.PaginationMeta            `json:"meta"`
		}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err == nil && cachedResult.Data != nil {
			return cachedResult.Data, cachedResult.Meta, nil
		}
	}

	fetchLimit := req.Limit * 2
	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("invalid url: %w", err)
	}

	query := reqURL.Query()
	query.Set("text", cleanSearch)
	query.Set("limit", fmt.Sprintf("%d", fetchLimit))
	query.Set("offset", fmt.Sprintf("%d", offset))

	// CATATAN: filter countrycode DIHAPUS agar pencarian global/general (seperti Baker Street) berfungsi.
	// Jika ingin mengutamakan Indonesia tanpa memblokir negara lain, Geoapify tidak perlu 'filter=countrycode:id'.

	query.Set("lang", "id")
	query.Set("apiKey", c.cfg.GeoapifyAPIKey)
	reqURL.RawQuery = query.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("failed to call geoapify service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return items, api.PaginationMeta{}, fmt.Errorf("geoapify error (%d): %s", resp.StatusCode, string(body))
	}

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("failed to decode response payload: %w", err)
	}

	for _, f := range geoResp.Features {
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

		case "district":
			entityType = "district"
			title = p.District
			if title == "" {
				title = p.Suburb
			}
			if cityClean != "" {
				subtitleParts = append(subtitleParts, cityClean)
			}
			if stateClean != "" {
				subtitleParts = append(subtitleParts, stateClean)
			}
			subtitleParts = append(subtitleParts, country)

		case "suburb":
			entityType = "suburb"
			title = p.Suburb
			if title == "" {
				title = p.District
			}
			if p.District != "" && p.District != title {
				subtitleParts = append(subtitleParts, p.District)
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
			if p.District != "" && p.District != p.Suburb {
				subtitleParts = append(subtitleParts, p.District)
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

		if title == "" {
			continue
		}

		items = append(items, dto.AreaGeneralResponseData{
			Type:        entityType,
			Title:       title,
			Subtitle:    subtitle,
			FullLabel:   fullLabel,
			Country:     country,
			CountryCode: strings.ToUpper(p.CountryCode),
			City:        cityClean,
			Province:    stateClean,
			District:    p.District,
			Suburb:      p.Suburb,
		})
	}

	hasNextPage := false
	if len(items) > req.Limit {
		hasNextPage = true
		items = items[:req.Limit]
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
// 2. SEARCH DETAIL
// -----------------------------------------------------------------------------
func (c *GeoapifyClient) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	cleanSearch := strings.TrimSpace(req.Search)

	items := make([]dto.AreaDetailResponseData, 0)

	if len(cleanSearch) < 2 {
		return items, api.PaginationMeta{
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: false,
		}, nil
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit

	cacheKey := fmt.Sprintf("geoapify:detail_v3:%s:%d:%d", cleanSearch, req.Page, req.Limit)
	cachedData, err := c.rdb.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedResult struct {
			Data []dto.AreaDetailResponseData `json:"data"`
			Meta api.PaginationMeta           `json:"meta"`
		}
		if err := json.Unmarshal([]byte(cachedData), &cachedResult); err == nil && cachedResult.Data != nil {
			return cachedResult.Data, cachedResult.Meta, nil
		}
	}

	fetchLimit := req.Limit * 2
	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("invalid url: %w", err)
	}

	query := reqURL.Query()
	query.Set("text", cleanSearch)
	query.Set("limit", fmt.Sprintf("%d", fetchLimit))
	query.Set("offset", fmt.Sprintf("%d", offset))
	query.Set("lang", "id")
	query.Set("apiKey", c.cfg.GeoapifyAPIKey)
	reqURL.RawQuery = query.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("failed to call geoapify service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return items, api.PaginationMeta{}, fmt.Errorf("geoapify error (%d): %s", resp.StatusCode, string(body))
	}

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return items, api.PaginationMeta{}, fmt.Errorf("failed to decode response payload: %w", err)
	}

	for _, f := range geoResp.Features {
		p := f.Properties

		street := p.Street
		if street == "" {
			street = p.Name
		}

		country := p.Country
		if country == "" {
			country = "Indonesia"
		}

		items = append(items, dto.AreaDetailResponseData{
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
		})
	}

	hasNextPage := false
	if len(items) > req.Limit {
		hasNextPage = true
		items = items[:req.Limit]
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
