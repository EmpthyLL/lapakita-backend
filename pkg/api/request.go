package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type BasePaginationRequest struct {
	Page        int      `form:"page" binding:"omitempty,min=1"`
	Limit       int      `form:"limit" binding:"omitempty,min=1"`
	SelectedID  string   `form:"selected_id" binding:"omitempty"`
	SelectedIDs []string `form:"selected_ids" binding:"omitempty,dive,required"`
}

// GetSelectedIDs combines the legacy single selection and multi-selection
// values, preserving request order while dropping empty values and duplicates.
func (r BasePaginationRequest) GetSelectedIDs() []string {
	values := make([]string, 0, len(r.SelectedIDs)+1)
	if r.SelectedID != "" {
		values = append(values, r.SelectedID)
	}
	values = append(values, r.SelectedIDs...)

	seen := make(map[string]struct{}, len(values))
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func (r *BasePaginationRequest) SetDefaults() {
	if r.Page <= 0 {
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
