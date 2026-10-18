package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type BasePaginationRequest struct {
	Page       int    `form:"page" binding:"omitempty,min=1"`
	Limit      int    `form:"limit" binding:"omitempty,min=1"`
	SelectedID string `form:"selected_id" binding:"omitempty"`
}

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
