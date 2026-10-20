package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type BasePaginationRequest struct {
	Page       int    `form:"page,default=1" binding:"omitempty"`
	Limit      int    `form:"limit,default=10" binding:"omitempty,min=1"`
	SelectedID string `form:"selected_id" binding:"omitempty"`
}

// GetSelectedIDs menguraikan query string koma (misal: "id1,id2,id3")
// menjadi []string bersih tanpa duplikat.
func (r BasePaginationRequest) GetSelectedIDs() []string {
	if r.SelectedID == "" {
		return nil
	}

	rawItems := strings.Split(r.SelectedID, ",")
	seen := make(map[string]struct{}, len(rawItems))
	result := make([]string, 0, len(rawItems))

	for _, item := range rawItems {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if _, exists := seen[item]; exists {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// GetPrimaryAnchor mengambil titik acuan tengah dari daftar SelectedID
func (r BasePaginationRequest) GetPrimaryAnchor() (string, bool) {
	ids := r.GetSelectedIDs()
	n := len(ids)
	if n == 0 {
		return "", false
	}

	var centerIdx int
	if n%2 == 0 {
		centerIdx = (n / 2) - 1
	} else {
		centerIdx = n / 2
	}

	return ids[centerIdx], true
}

func (r *BasePaginationRequest) SetDefaults() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit <= 0 {
		r.Limit = 10
	}
}

func GetLanguageFromHeader(c *gin.Context) string {
	lang := c.GetHeader("lang")
	if lang == "" {
		lang = "en"
	}
	return lang
}
