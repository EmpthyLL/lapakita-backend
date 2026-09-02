package dto

import "lapakita-backend/pkg/api"

type GetOwnerStallsRequest struct {
	api.BasePaginationRequest

	// Required Owner Identifier
	OwnerID string `form:"owner_id" binding:"required"`

	// Specific Search Filters
	Title          string `form:"title"`           // Search by Title
	PropertyType   string `form:"property_type"`   // Filter by Property Type
	PermanenceType string `form:"permanence_type"` // Filter by Permanence Type
	Placement      string `form:"placement"`       // Filter by Placement
	Location       string `form:"location"`        // Search address (street, suburb, district, city, province)
	IsPublished    *bool  `form:"is_published"`    // Filter status publish (nil = all)
}
