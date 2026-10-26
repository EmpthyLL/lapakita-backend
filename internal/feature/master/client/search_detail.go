package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
)

// -----------------------------------------------------------------------------
// 2. SEARCH DETAIL (Form Auto-fill Owner)
// -----------------------------------------------------------------------------
func (c *MasterClient) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	cleanSearch := strings.TrimSpace(req.Search)
	items := make([]dto.AreaDetailResponseData, 0)

	metaFallback := api.PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		HasNextPage: false,
		HasPrevPage: page > 1,
	}

	if len(cleanSearch) < 2 {
		return items, metaFallback, nil
	}

	offset := (page - 1) * limit

	cacheKey := fmt.Sprintf("geoapify:det_v9:%s:%d:%d", cleanSearch, page, limit)
	if c.rdb != nil {
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
	}

	apiURL := "https://api.geoapify.com/v1/geocode/autocomplete"
	reqURL, err := url.Parse(apiURL)
	if err != nil {
		return items, metaFallback, nil
	}

	query := reqURL.Query()
	query.Set("text", cleanSearch)
	query.Set("limit", fmt.Sprintf("%d", limit*2))
	query.Set("offset", fmt.Sprintf("%d", offset))
	// query.Set("lang", "en")
	query.Set("lang", "id")
	query.Set("countrycodes", "id")
	query.Set("apiKey", c.cfg.GeoapifyAPIKey)
	reqURL.RawQuery = query.Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return items, metaFallback, nil
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil || resp == nil {
		return items, metaFallback, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return items, metaFallback, nil
	}

	var geoResp dto.GeoapifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoResp); err != nil {
		return items, metaFallback, nil
	}

	for _, f := range geoResp.Features {
		p := f.Properties
		countryCode := strings.ToUpper(p.CountryCode)

		street := p.Street
		if street == "" {
			street = p.Name
		}
		if street == "" {
			street = p.Formatted
		}

		country := p.Country

		items = append(items, dto.AreaDetailResponseData{
			Formatted:      p.Formatted,
			StreetAddress:  street,
			Suburb:         p.Suburb,
			District:       p.District,
			City:           translateIndonesianDirection(cleanComponent(p.City, countryCode), countryCode),
			Province:       translateIndonesianDirection(cleanComponent(p.State, countryCode), countryCode),
			Country:        country,
			CountryCode:    countryCode,
			PostalCode:     p.Postcode,
			Latitude:       p.Lat,
			Longitude:      p.Lon,
			MapURL:         fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%f,%f", p.Lat, p.Lon),
			EmbeddedMapURL: fmt.Sprintf("https://maps.google.com/maps?q=%f,%f&hl=id&z=16&output=embed", p.Lat, p.Lon),
		})
	}

	hasNextPage := false
	if len(items) > limit {
		hasNextPage = true
		items = items[:limit]
	}

	meta := api.PaginationMeta{
		CurrentPage: page,
		PerPage:     limit,
		HasNextPage: hasNextPage,
		HasPrevPage: page > 1,
	}

	if c.rdb != nil {
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
	}

	return items, meta, nil
}
