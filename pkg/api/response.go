package api

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginationMeta struct {
	TotalItems  int  `json:"totalItems"`
	TotalPages  int  `json:"totalPages"`
	CurrentPage int  `json:"currentPage"`
	PerPage     int  `json:"perPage"`
	HasNextPage bool `json:"hasNextPage"`
	HasPrevPage bool `json:"hasPrevPage"`
}

type PageResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(c *gin.Context, statusCode int, message string, data interface{}, meta PaginationMeta) {
	c.JSON(statusCode, PageResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, errs interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Errors:  errs,
	})
}
