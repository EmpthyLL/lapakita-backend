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
// 1. SEARCH GENERAL (Dropdown UI Search Bar)
// -----------------------------------------------------------------------------
func (c *MasterClient) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}

	cleanSearch := strings.TrimSpace(req.Search)
	items := make([]dto.AreaGeneralResponseData, 0)

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

	// Cache Key
	cacheKey := fmt.Sprintf("geoapify:gen_v9:%s:%d:%d", cleanSearch, page, limit)
	if c.rdb != nil {
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

		cityClean := translateIndonesianDirection(cleanComponent(p.City, countryCode), countryCode)
		stateClean := translateIndonesianDirection(cleanComponent(p.State, countryCode), countryCode)
		country := p.Country

		entityType := "street"
		title := p.Street
		if title == "" {
			title = p.Name
		}
		if title == "" {
			title = p.Formatted
		}

		if title == "" {
			continue
		}

		title = translateIndonesianDirection(cleanComponent(title, countryCode), countryCode)

		var subtitleParts []string

		switch p.ResultType {
		case "country":
			entityType = "country"
			title = country
			subtitleParts = []string{"Country"}

		case "state", "province":
			entityType = "province"
			title = stateClean
			if country != "" {
				subtitleParts = []string{country}
			}

		case "city", "county":
			entityType = "city"
			title = cityClean
			if stateClean != "" {
				subtitleParts = append(subtitleParts, stateClean)
			}
			if country != "" {
				subtitleParts = append(subtitleParts, country)
			}

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
			if country != "" {
				subtitleParts = append(subtitleParts, country)
			}

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
			if country != "" {
				subtitleParts = append(subtitleParts, country)
			}

		default:
			entityType = "street"
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
			if country != "" {
				subtitleParts = append(subtitleParts, country)
			}
		}

		subtitle := strings.Join(subtitleParts, ", ")
		fullLabel := title
		if subtitle != "" && entityType != "country" {
			fullLabel = fmt.Sprintf("%s, %s", title, subtitle)
		}

		items = append(items, dto.AreaGeneralResponseData{
			Type:        entityType,
			Title:       title,
			Subtitle:    subtitle,
			FullLabel:   fullLabel,
			Country:     country,
			CountryCode: countryCode,
			City:        cityClean,
			Province:    stateClean,
			District:    p.District,
			Suburb:      p.Suburb,
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
			Data []dto.AreaGeneralResponseData `json:"data"`
			Meta api.PaginationMeta            `json:"meta"`
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
