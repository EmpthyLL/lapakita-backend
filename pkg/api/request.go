package api

import "github.com/gin-gonic/gin"

type BasePaginationRequest struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1"`
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
