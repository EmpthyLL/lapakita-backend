package database

import (
	"fmt"
	"math"
	"strings"

	"lapakita-backend/pkg/api"

	"gorm.io/gorm"
)

// -----------------------------------------------------------------------------
// 1. AUTOPAGINATE UNTUK QUERY GORM (SQL DATABASE)
// -----------------------------------------------------------------------------

func AutoPaginate[T any](
	db *gorm.DB,
	req api.BasePaginationRequest,
	targetTableName string,
	outSlice *[]T,
	meta *api.PaginationMeta,
	customOrderOpsional string,
	primaryKeyColumn ...string,
) error {
	req.SetDefaults()

	pk := "id"
	if len(primaryKeyColumn) > 0 && primaryKeyColumn[0] != "" {
		pk = primaryKeyColumn[0]
	}

	orderCol := "created_at"
	orderDir := "DESC"

	if customOrderOpsional != "" {
		orderCol = customOrderOpsional
		if len(customOrderOpsional) > 4 && customOrderOpsional[len(customOrderOpsional)-4:] == " ASC" {
			orderDir = "ASC"
			orderCol = strings.TrimSpace(customOrderOpsional[:len(customOrderOpsional)-4])
		} else if len(customOrderOpsional) > 5 && customOrderOpsional[len(customOrderOpsional)-5:] == " DESC" {
			orderDir = "DESC"
			orderCol = strings.TrimSpace(customOrderOpsional[:len(customOrderOpsional)-5])
		}
	}

	var totalItems int64
	countQuery := db.Session(&gorm.Session{})
	if err := countQuery.Count(&totalItems).Error; err != nil {
		return err
	}

	var hasNextPage, hasPrevPage bool

	// -------------------------------------------------------------------------
	// KASUS A: SelectedID Active (Anchor Center Hydration)
	// -------------------------------------------------------------------------
	if req.SelectedID != "" {
		var anchor T
		err := db.Session(&gorm.Session{}).
			Where(fmt.Sprintf("%s = ?", pk), req.SelectedID).
			First(&anchor).Error

		if err == nil {
			beforeLimit := 4
			if req.Limit < 5 {
				beforeLimit = req.Limit / 2
			}
			afterLimit := req.Limit - beforeLimit

			anchorValSubquery := fmt.Sprintf("(SELECT %s FROM %s WHERE %s = ? LIMIT 1)", orderCol, targetTableName, pk)

			var upperOp, upperOrder, lowerOp, lowerOrder string
			if orderDir == "ASC" {
				upperOp = fmt.Sprintf("(%s < %s OR (%s = %s AND %s < ?))", orderCol, anchorValSubquery, orderCol, anchorValSubquery, pk)
				upperOrder = fmt.Sprintf("%s DESC, %s DESC", orderCol, pk)
				lowerOp = fmt.Sprintf("(%s > %s OR (%s = %s AND %s >= ?))", orderCol, anchorValSubquery, orderCol, anchorValSubquery, pk)
				lowerOrder = fmt.Sprintf("%s ASC, %s ASC", orderCol, pk)
			} else {
				upperOp = fmt.Sprintf("(%s > %s OR (%s = %s AND %s > ?))", orderCol, anchorValSubquery, orderCol, anchorValSubquery, pk)
				upperOrder = fmt.Sprintf("%s ASC, %s ASC", orderCol, pk)
				lowerOp = fmt.Sprintf("(%s < %s OR (%s = %s AND %s <= ?))", orderCol, anchorValSubquery, orderCol, anchorValSubquery, pk)
				lowerOrder = fmt.Sprintf("%s DESC, %s DESC", orderCol, pk)
			}

			// 1. Fetch item SEBELUM anchor (Upper)
			var upperSlice []T
			err = db.Session(&gorm.Session{}).
				Where(upperOp, req.SelectedID, req.SelectedID).
				Order(upperOrder).
				Limit(beforeLimit).
				Find(&upperSlice).Error
			if err != nil {
				return err
			}

			// Reverse Upper Slice
			for i, j := 0, len(upperSlice)-1; i < j; i, j = i+1, j-1 {
				upperSlice[i], upperSlice[j] = upperSlice[j], upperSlice[i]
			}

			// 2. Fetch item MULAI DARI anchor KE BAWAH (Lower)
			var lowerSlice []T
			err = db.Session(&gorm.Session{}).
				Where(lowerOp, req.SelectedID, req.SelectedID).
				Order(lowerOrder).
				Limit(afterLimit).
				Find(&lowerSlice).Error
			if err != nil {
				return err
			}

			combined := append(upperSlice, lowerSlice...)
			*outSlice = combined

			if len(combined) > 0 {
				var prevCount int64
				_ = db.Session(&gorm.Session{}).
					Where(upperOp, req.SelectedID, req.SelectedID).
					Count(&prevCount).Error
				hasPrevPage = prevCount > 0
				hasNextPage = int64(len(lowerSlice)) >= int64(afterLimit)
			}

			totalPages := int((totalItems + int64(req.Limit) - 1) / int64(req.Limit))

			*meta = api.PaginationMeta{
				TotalItems:  int(totalItems),
				TotalPages:  totalPages,
				CurrentPage: req.Page,
				PerPage:     req.Limit,
				HasNextPage: hasNextPage,
				HasPrevPage: hasPrevPage,
			}

			return nil
		}
	}

	// -------------------------------------------------------------------------
	// KASUS B: Standard Offset Pagination (Fallback)
	// -------------------------------------------------------------------------
	offset := (req.Page - 1) * req.Limit
	orderClause := fmt.Sprintf("%s %s, %s %s", orderCol, orderDir, pk, orderDir)

	err := db.Session(&gorm.Session{}).
		Order(orderClause).
		Limit(req.Limit + 1).
		Offset(offset).
		Find(outSlice).Error
	if err != nil {
		return err
	}

	items := *outSlice
	if len(items) > req.Limit {
		hasNextPage = true
		*outSlice = items[:req.Limit]
	}
	hasPrevPage = req.Page > 1

	totalPages := int((totalItems + int64(req.Limit) - 1) / int64(req.Limit))

	*meta = api.PaginationMeta{
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
	}

	return nil
}

// -----------------------------------------------------------------------------
// 2. AUTOPAGINATESLICE UNTUK IN-MEMORY ARRAYS / JSONB SLICES
// -----------------------------------------------------------------------------

func AutoPaginateSlice[T any](
	sourceSlice []T,
	req api.BasePaginationRequest,
	outSlice *[]T,
	meta *api.PaginationMeta,
	matchSelected func(item T, selectedID string) bool,
) {
	req.SetDefaults()

	totalItems := len(sourceSlice)
	if totalItems == 0 {
		*outSlice = []T{}
		*meta = api.PaginationMeta{
			TotalItems:  0,
			TotalPages:  0,
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: false,
		}
		return
	}

	// -------------------------------------------------------------------------
	// KASUS A: SelectedID Active (Hydration Dropdown Tanpa Duplikasi)
	// -------------------------------------------------------------------------
	if req.SelectedID != "" && matchSelected != nil {
		selectedIdx := -1
		for idx, item := range sourceSlice {
			if matchSelected(item, req.SelectedID) {
				selectedIdx = idx
				break
			}
		}

		if selectedIdx != -1 {
			// The selected option is an extra hydration item on the first page.
			// Remove it from the regular stream so it cannot appear again later.
			regularItems := make([]T, 0, totalItems-1)
			regularItems = append(regularItems, sourceSlice[:selectedIdx]...)
			regularItems = append(regularItems, sourceSlice[selectedIdx+1:]...)

			startIdx := 0
			if req.Page > 1 {
				startIdx = (req.Page-1)*req.Limit - 1
				if startIdx < 0 {
					startIdx = 0
				}
			}
			if startIdx >= len(regularItems) && req.Page > 1 {
				*outSlice = []T{}
				*meta = api.PaginationMeta{
					TotalItems:  totalItems,
					TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
					CurrentPage: req.Page,
					PerPage:     req.Limit,
					HasPrevPage: true,
				}
				return
			}

			pageSize := req.Limit
			if req.Page == 1 {
				pageSize--
			}
			endIdx := startIdx + pageSize
			if endIdx > len(regularItems) {
				endIdx = len(regularItems)
			}

			pagedItems := regularItems[startIdx:endIdx]
			if req.Page == 1 {
				pagedItems = append([]T{sourceSlice[selectedIdx]}, pagedItems...)
			}
			hasPrevPage := req.Page > 1
			hasNextPage := endIdx < len(regularItems)
			totalPages := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

			*outSlice = pagedItems
			*meta = api.PaginationMeta{
				TotalItems:  totalItems,
				TotalPages:  totalPages,
				CurrentPage: req.Page,
				PerPage:     req.Limit,
				HasNextPage: hasNextPage,
				HasPrevPage: hasPrevPage,
			}
			return
		}
	}

	// -------------------------------------------------------------------------
	// KASUS B: Standard Offset Slicing
	// -------------------------------------------------------------------------
	startIdx := (req.Page - 1) * req.Limit
	if startIdx >= totalItems {
		*outSlice = []T{}
		*meta = api.PaginationMeta{
			TotalItems:  totalItems,
			TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: false,
			HasPrevPage: req.Page > 1,
		}
		return
	}

	endIdx := startIdx + req.Limit
	if endIdx > totalItems {
		endIdx = totalItems
	}

	pagedItems := sourceSlice[startIdx:endIdx]
	hasPrevPage := req.Page > 1
	hasNextPage := endIdx < totalItems
	totalPages := int(math.Ceil(float64(totalItems) / float64(req.Limit)))

	*outSlice = pagedItems
	*meta = api.PaginationMeta{
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: hasNextPage,
		HasPrevPage: hasPrevPage,
	}
}
