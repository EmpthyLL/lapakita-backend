package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type BasePaginationRequest struct {
	Page        int      `form:"page" binding:"omitempty,min=1"`
	Limit       int      `form:"limit" binding:"omitempty,min=1"`
	SelectedIDs []string `form:"selected_id" binding:"omitempty,dive,required"`
}

func (r BasePaginationRequest) GetSelectedIDs() []string {
	seen := make(map[string]struct{}, len(r.SelectedIDs))
	result := make([]string, 0, len(r.SelectedIDs))
	for _, selection := range r.SelectedIDs {
		for _, value := range strings.Split(selection, ",") {
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
