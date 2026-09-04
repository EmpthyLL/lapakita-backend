package client

import (
	"context"
	"encoding/json"
	"fmt"
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

// cleanComponent membersihkan prefix administrative baik bahasa Indonesia maupun Inggris
// cleanComponent membersihkan prefix & suffix administrative baik bahasa Indonesia maupun Inggris
func cleanComponent(val string, countryCode string) string {
	cleaned := strings.TrimSpace(val)

	// Prefix Bahasa Inggris (Geoapify Output)
	cleaned = strings.TrimPrefix(cleaned, "City of ")
	cleaned = strings.TrimPrefix(cleaned, "Regency of ")
	cleaned = strings.TrimPrefix(cleaned, "Province of ")
	cleaned = strings.TrimPrefix(cleaned, "Special Region of ")
	cleaned = strings.TrimPrefix(cleaned, "Special Capital Region of ")

	// Khusus Indonesia (CountryCode == "ID")
	if strings.ToUpper(countryCode) == "ID" {
		// Hapus Suffix Bahasa Inggris yang sering ditempelkan Geoapify di akhir nama kota/kabupaten
		cleaned = strings.TrimSuffix(cleaned, " City")
		cleaned = strings.TrimSuffix(cleaned, " Regency")
	}

	// Prefix Bahasa Indonesia
	cleaned = strings.TrimPrefix(cleaned, "Kota ")
	cleaned = strings.TrimPrefix(cleaned, "Kabupaten ")
	cleaned = strings.TrimPrefix(cleaned, "Provinsi ")
	cleaned = strings.TrimPrefix(cleaned, "Propinsi ")
	cleaned = strings.TrimPrefix(cleaned, "Daerah Khusus Ibukota ")
	cleaned = strings.TrimPrefix(cleaned, "Daerah Istimewa ")

	return strings.TrimSpace(cleaned)
}

// translateIndonesianDirection memutar posisi & menerjemahkan arah angin HANYA untuk Indonesia (CountryCode == "ID")
func translateIndonesianDirection(val string, countryCode string) string {
	if strings.ToUpper(countryCode) != "ID" || val == "" {
		return val // Luar negeri tidak disentuh (aman dari risiko salah terjemah)
	}

	// Map arah angin dari Depan (Inggris) -> Belakang (Indonesia)
	directions := map[string]string{
		"North ":     " Utara",
		"South ":     " Selatan",
		"West ":      " Barat",
		"East ":      " Timur",
		"Central ":   " Tengah",
		"Southeast ": " Tenggara",
		"Southwest ": " Barat Daya",
		"Northeast ": " Timur Laut",
		"Northwest ": " Barat Daya",
	}

	for engPrefix, idSuffix := range directions {
		if strings.HasPrefix(val, engPrefix) {
			baseName := strings.TrimPrefix(val, engPrefix)
			return baseName + idSuffix
		}
	}

	return val
}

// -----------------------------------------------------------------------------
// 1. SEARCH GENERAL (Dropdown UI Search Bar)
// -----------------------------------------------------------------------------
func (c *GeoapifyClient) SearchGeneral(ctx context.Context, req dto.GetAreaGeneralRequest) ([]dto.AreaGeneralResponseData, api.PaginationMeta, error) {
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
	query.Set("lang", "en")
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

// -----------------------------------------------------------------------------
// 2. SEARCH DETAIL (Form Auto-fill Owner)
// -----------------------------------------------------------------------------
func (c *GeoapifyClient) SearchDetail(ctx context.Context, req dto.GetAreaDetailRequest) ([]dto.AreaDetailResponseData, api.PaginationMeta, error) {
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
	query.Set("lang", "en")
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
