package database

import (
	"fmt"
	"math"
	"reflect"
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
	selectedIDs := req.GetSelectedIDs()
	anchorID := ""
	if len(selectedIDs) > 0 {
		anchorID = normalizeSelectedID(selectedIDs[0])
	}

	// -------------------------------------------------------------------------
	// KASUS A: SelectedID Active (Anchor Center Hydration)
	// -------------------------------------------------------------------------
	if anchorID != "" {
		var anchor T
		err := db.Session(&gorm.Session{}).
			Where(fmt.Sprintf("%s = ?", pk), anchorID).
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
				Where(upperOp, anchorID, anchorID).
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
				Where(lowerOp, anchorID, anchorID).
				Order(lowerOrder).
				Limit(afterLimit).
				Find(&lowerSlice).Error
			if err != nil {
				return err
			}

			combined := append(upperSlice, lowerSlice...)

			if len(combined) > 0 {
				var prevCount int64
				_ = db.Session(&gorm.Session{}).
					Where(upperOp, anchorID, anchorID).
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

			combined, err = hydrateSelectedRows(db, pk, selectedIDs, combined)
			if err != nil {
				return err
			}
			*outSlice = combined

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
	hydratedItems, hydrateErr := hydrateSelectedRows(db, pk, selectedIDs, *outSlice)
	if hydrateErr != nil {
		return hydrateErr
	}
	*outSlice = hydratedItems
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

func normalizeSelectedID(value string) string {
	if normalized, err := api.Normalize(value); err == nil {
		return normalized
	}
	return value
}

func entityID[T any](item T) string {
	value := reflect.ValueOf(item)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return ""
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return ""
	}
	field := value.FieldByName("ID")
	if !field.IsValid() || !field.CanInterface() {
		return ""
	}
	return normalizeSelectedID(fmt.Sprint(field.Interface()))
}

func hydrateSelectedRows[T any](db *gorm.DB, primaryKey string, selectedIDs []string, items []T) ([]T, error) {
	if len(selectedIDs) == 0 {
		return items, nil
	}

	queryIDs := make([]string, 0, len(selectedIDs))
	for _, id := range selectedIDs {
		queryIDs = append(queryIDs, normalizeSelectedID(id))
	}

	var selectedItems []T
	if err := db.Session(&gorm.Session{}).
		Where(fmt.Sprintf("%s IN ?", primaryKey), queryIDs).
		Find(&selectedItems).Error; err != nil {
		return nil, err
	}

	selectedByID := make(map[string]T, len(selectedItems))
	for _, item := range selectedItems {
		selectedByID[entityID(item)] = item
	}
	seen := make(map[string]struct{}, len(items))
	for _, item := range items {
		seen[entityID(item)] = struct{}{}
	}

	result := make([]T, 0, len(selectedItems)+len(items))
	for _, id := range queryIDs {
		if _, alreadyIncluded := seen[id]; alreadyIncluded {
			continue
		}
		if item, exists := selectedByID[id]; exists {
			result = append(result, item)
			seen[id] = struct{}{}
		}
	}
	return append(result, items...), nil
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
	selectedIDs := req.GetSelectedIDs()
	if len(selectedIDs) > 0 && matchSelected != nil {
		selectedItems := make([]T, 0, len(selectedIDs))
		selectedIndexes := make(map[int]struct{}, len(selectedIDs))
		for _, selectedID := range selectedIDs {
			for index, item := range sourceSlice {
				if matchSelected(item, selectedID) {
					if _, alreadySelected := selectedIndexes[index]; !alreadySelected {
						selectedItems = append(selectedItems, item)
						selectedIndexes[index] = struct{}{}
					}
					break
				}
			}
		}

		regularItems := make([]T, 0, totalItems-len(selectedItems))
		for index, item := range sourceSlice {
			if _, isSelected := selectedIndexes[index]; !isSelected {
				regularItems = append(regularItems, item)
			}
		}

		startIdx := 0
		pageSize := req.Limit - len(selectedItems)
		if req.Page > 1 {
			startIdx = (req.Page-1)*req.Limit - len(selectedItems)
			pageSize = req.Limit
		}
		if pageSize < 0 {
			pageSize = 0
		}
		if startIdx < 0 {
			startIdx = 0
		}
		if startIdx >= len(regularItems) {
			*outSlice = []T{}
			*meta = api.PaginationMeta{
				TotalItems:  totalItems,
				TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
				CurrentPage: req.Page,
				PerPage:     req.Limit,
				HasPrevPage: req.Page > 1,
			}
			return
		}
		endIdx := startIdx + pageSize
		if endIdx > len(regularItems) {
			endIdx = len(regularItems)
		}
		pagedItems := regularItems[startIdx:endIdx]
		if req.Page == 1 {
			pagedItems = append(selectedItems, pagedItems...)
		}
		*outSlice = pagedItems
		*meta = api.PaginationMeta{
			TotalItems:  totalItems,
			TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
			CurrentPage: req.Page,
			PerPage:     req.Limit,
			HasNextPage: endIdx < len(regularItems),
			HasPrevPage: req.Page > 1,
		}
		return
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
