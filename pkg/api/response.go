package api

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Standar Response Sukses / Generik
type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// Standar Metadata Pagination
type PaginationMeta struct {
	TotalItems  int  `json:"totalItems,omitempty"`
	TotalPages  int  `json:"totalPages,omitempty"`
	CurrentPage int  `json:"currentPage"`
	PerPage     int  `json:"perPage"`
	HasNextPage bool `json:"hasNextPage"`
	HasPrevPage bool `json:"hasPrevPage"`
}

// Standar Response Khusus Pagination (Generic)
type PaginatedResponse[T any] struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    T              `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

// Struct Khusus Detail Error Per-Field (Validation Error)
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Standar Response Error
type ErrorResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors,omitempty"`
}

// --- Helper Functions ---

// Success mengirim respon sukses tanpa pagination
func Success[T any](c *gin.Context, status int, message string, data T) {
	c.Header("Content-Type", "application/json")
	c.Status(status)
	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(Response[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessWithPagination mengirim respon data ber-halaman
func SuccessWithPagination[T any](c *gin.Context, status int, message string, data T, meta PaginationMeta) {
	c.Header("Content-Type", "application/json")
	c.Status(status)
	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(PaginatedResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Error mengirim respon error umum
func Error(c *gin.Context, status int, message string) {
	ErrorWithFields(c, status, message, nil)
}

// ErrorWithFields mengirim respon error beserta detail validation error per field
func ErrorWithFields(c *gin.Context, status int, message string, fields []FieldError) {
	c.Header("Content-Type", "application/json")
	c.Status(status)
	encoder := json.NewEncoder(c.Writer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(ErrorResponse{
		Success: false,
		Message: message,
		Errors:  fields,
	})
}
