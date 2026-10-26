package client

import (
	"net/http"
	"strings"
	"time"

	"lapakita-backend/config"

	"github.com/redis/go-redis/v9"
)

type MasterClient struct {
	cfg        *config.Config
	rdb        *redis.Client
	httpClient *http.Client
}

func NewMasterClient(cfg *config.Config, rdb *redis.Client) *MasterClient {
	return &MasterClient{
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
