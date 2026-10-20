package database

import (
	"fmt"
	"math"
	"reflect"
	"strings"

	"lapakita-backend/pkg/api"

	"gorm.io/gorm"
)

// Helper presisi untuk mengekstrak string UUID/ID murni dari struct/pointer
func getItemID[T any](item T, pkColumn string) string {
	val := reflect.ValueOf(item)
	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return ""
		}
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		field := val.FieldByName("ID")
		if field.IsValid() && field.CanInterface() {
			return normalizeSelectedID(fmt.Sprintf("%v", field.Interface()))
		}

		for i := 0; i < val.NumField(); i++ {
			typeField := val.Type().Field(i)
			tag := typeField.Tag.Get("gorm")
			if strings.Contains(tag, "primaryKey") || strings.EqualFold(typeField.Name, pkColumn) {
				return normalizeSelectedID(fmt.Sprintf("%v", val.Field(i).Interface()))
			}
		}
	}

	return normalizeSelectedID(fmt.Sprintf("%v", item))
}

func normalizeSelectedID(value string) string {
	if normalized, err := api.Normalize(value); err == nil {
		return normalized
	}
	return value
}

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
	pk := "id"
	if len(primaryKeyColumn) > 0 && primaryKeyColumn[0] != "" {
		pk = primaryKeyColumn[0]
	}

	orderCol := "created_at"
	orderDir := "DESC"

	if customOrderOpsional != "" {
		orderCol = customOrderOpsional
		if len(customOrderOpsional) > 4 && strings.ToUpper(customOrderOpsional[len(customOrderOpsional)-4:]) == " ASC" {
			orderDir = "ASC"
			orderCol = strings.TrimSpace(customOrderOpsional[:len(customOrderOpsional)-4])
		} else if len(customOrderOpsional) > 5 && strings.ToUpper(customOrderOpsional[len(customOrderOpsional)-5:]) == " DESC" {
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

	// -------------------------------------------------------------------------
	// KASUS A: SelectedID Active (Anchor Center Hydration & Bidirectional Page)
	// -------------------------------------------------------------------------
	if anchorID, ok := req.GetPrimaryAnchor(); ok {
		var anchor T
		err := db.Session(&gorm.Session{}).
			Where(fmt.Sprintf("%s = ?", pk), anchorID).
			First(&anchor).Error

		if err == nil {
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

			var upperSlice []T
			var lowerSlice []T

			if req.Page == 1 {
				idealUpperLimit := 4
				if req.Limit < 10 {
					idealUpperLimit = req.Limit / 2
				}

				// Fetch Upper (Item sebelum Anchor)
				err = db.Session(&gorm.Session{}).
					Where(upperOp, anchorID, anchorID, anchorID).
					Order(upperOrder).
					Limit(idealUpperLimit).
					Find(&upperSlice).Error
				if err != nil {
					return err
				}

				actualUpperCount := len(upperSlice)
				lowerLimit := req.Limit - actualUpperCount

				// Fetch Lower (Item mulai dari Anchor ke bawah)
				err = db.Session(&gorm.Session{}).
					Where(lowerOp, anchorID, anchorID, anchorID).
					Order(lowerOrder).
					Limit(lowerLimit + 1).
					Find(&lowerSlice).Error
				if err != nil {
					return err
				}

				if len(lowerSlice) > lowerLimit {
					hasNextPage = true
					lowerSlice = lowerSlice[:lowerLimit]
				}

				actualLowerCount := len(lowerSlice)
				if (actualUpperCount + actualLowerCount) < req.Limit {
					extraNeeded := req.Limit - (actualUpperCount + actualLowerCount)
					var extraUpper []T
					_ = db.Session(&gorm.Session{}).
						Where(upperOp, anchorID, anchorID, anchorID).
						Order(upperOrder).
						Offset(actualUpperCount).
						Limit(extraNeeded).
						Find(&extraUpper).Error
					if len(extraUpper) > 0 {
						upperSlice = append(upperSlice, extraUpper...)
					}
				}

				var totalUpper int64
				_ = db.Session(&gorm.Session{}).
					Where(upperOp, anchorID, anchorID, anchorID).
					Count(&totalUpper).Error
				hasPrevPage = totalUpper > int64(len(upperSlice))

			} else if req.Page < 1 {
				// Halaman sebelum Anchor (Page 0, -1, -2, dst)
				pageOffset := (1 - req.Page) * req.Limit
				upperOffset := 4 + (pageOffset - req.Limit)

				err = db.Session(&gorm.Session{}).
					Where(upperOp, anchorID, anchorID, anchorID).
					Order(upperOrder).
					Offset(upperOffset).
					Limit(req.Limit + 1).
					Find(&upperSlice).Error
				if err != nil {
					return err
				}

				if len(upperSlice) > req.Limit {
					hasPrevPage = true
					upperSlice = upperSlice[:req.Limit]
				}
				hasNextPage = true

			} else {
				// Halaman setelah Anchor (Page 2, 3, 4, dst)
				pageOffset := (req.Page - 1) * req.Limit
				lowerOffset := (req.Limit - 4) + (pageOffset - req.Limit)

				err = db.Session(&gorm.Session{}).
					Where(lowerOp, anchorID, anchorID, anchorID).
					Order(lowerOrder).
					Offset(lowerOffset).
					Limit(req.Limit + 1).
					Find(&lowerSlice).Error
				if err != nil {
					return err
				}

				if len(lowerSlice) > req.Limit {
					hasNextPage = true
					lowerSlice = lowerSlice[:req.Limit]
				}
				hasPrevPage = true
			}

			// Reverse Upper Slice agar urutan kembali kronologis
			for i, j := 0, len(upperSlice)-1; i < j; i, j = i+1, j-1 {
				upperSlice[i], upperSlice[j] = upperSlice[j], upperSlice[i]
			}

			combined := append(upperSlice, lowerSlice...)

			// -----------------------------------------------------------------
			// BOUNDARY PINNING DINAMIS UNTUK ITEM TERPILIH LAINNYA
			// -----------------------------------------------------------------
			if len(selectedIDs) > 1 {
				var allSelectedEntities []T
				err = db.Session(&gorm.Session{}).
					Where(fmt.Sprintf("%s IN (?)", pk), selectedIDs).
					Order(fmt.Sprintf("%s %s, %s %s", orderCol, orderDir, pk, orderDir)).
					Find(&allSelectedEntities).Error

				if err == nil && len(allSelectedEntities) > 0 {
					existingMap := make(map[string]bool)
					for _, item := range combined {
						idStr := getItemID(item, pk)
						if idStr != "" {
							existingMap[idStr] = true
						}
					}

					var topPinned []T
					var bottomPinned []T

					for _, selItem := range allSelectedEntities {
						selID := getItemID(selItem, pk)
						if existingMap[selID] {
							continue
						}

						var checkUpper int64
						_ = db.Session(&gorm.Session{}).
							Where(fmt.Sprintf("%s = ? AND %s", pk, upperOp), selID, anchorID, anchorID, anchorID).
							Count(&checkUpper).Error

						if checkUpper > 0 {
							topPinned = append(topPinned, selItem)
						} else {
							bottomPinned = append(bottomPinned, selItem)
						}
					}

					finalCombined := make([]T, 0, len(topPinned)+len(combined)+len(bottomPinned))
					finalCombined = append(finalCombined, topPinned...)
					finalCombined = append(finalCombined, combined...)
					finalCombined = append(finalCombined, bottomPinned...)

					combined = finalCombined
				}
			}

			*outSlice = combined
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
	if offset < 0 {
		offset = 0
	}
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

	selectedIDs := req.GetSelectedIDs()

	if anchorID, ok := req.GetPrimaryAnchor(); ok && matchSelected != nil {
		anchorIdx := -1
		for idx, item := range sourceSlice {
			if matchSelected(item, anchorID) {
				anchorIdx = idx
				break
			}
		}

		if anchorIdx != -1 {
			beforeLimit := 4
			if req.Limit < 10 {
				beforeLimit = req.Limit / 2
			}

			var startIdx, endIdx int

			if req.Page == 1 {
				startIdx = anchorIdx - beforeLimit
				if startIdx < 0 {
					startIdx = 0
				}

				endIdx = startIdx + req.Limit
				if endIdx > totalItems {
					endIdx = totalItems
					startIdx = endIdx - req.Limit
					if startIdx < 0 {
						startIdx = 0
					}
				}
			} else if req.Page < 1 {
				pageOffset := (1 - req.Page) * req.Limit
				endIdx = anchorIdx - beforeLimit - (pageOffset - req.Limit)
				if endIdx < 0 {
					endIdx = 0
				}
				startIdx = endIdx - req.Limit
				if startIdx < 0 {
					startIdx = 0
				}
			} else {
				pageOffset := (req.Page - 1) * req.Limit
				startIdx = anchorIdx + (req.Limit - beforeLimit) + (pageOffset - req.Limit)
				if startIdx > totalItems {
					startIdx = totalItems
				}
				endIdx = startIdx + req.Limit
				if endIdx > totalItems {
					endIdx = totalItems
				}
			}

			pagedItems := sourceSlice[startIdx:endIdx]

			if len(selectedIDs) > 1 {
				existingMap := make(map[string]bool)
				for _, item := range pagedItems {
					idStr := getItemID(item, "id")
					if idStr != "" {
						existingMap[idStr] = true
					}
				}

				var topPinned []T
				var bottomPinned []T

				for idx, item := range sourceSlice {
					itemId := getItemID(item, "id")
					for _, selID := range selectedIDs {
						if itemId == normalizeSelectedID(selID) && !existingMap[itemId] {
							if idx < anchorIdx {
								topPinned = append(topPinned, item)
							} else {
								bottomPinned = append(bottomPinned, item)
							}
							existingMap[itemId] = true
						}
					}
				}

				finalItems := make([]T, 0, len(topPinned)+len(pagedItems)+len(bottomPinned))
				finalItems = append(finalItems, topPinned...)
				finalItems = append(finalItems, pagedItems...)
				finalItems = append(finalItems, bottomPinned...)

				pagedItems = finalItems
			}

			*outSlice = pagedItems
			*meta = api.PaginationMeta{
				TotalItems:  totalItems,
				TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
				CurrentPage: req.Page,
				PerPage:     req.Limit,
				HasNextPage: endIdx < totalItems,
				HasPrevPage: startIdx > 0,
			}
			return
		}
	}

	startIdx := (req.Page - 1) * req.Limit
	if startIdx < 0 {
		startIdx = 0
	}
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

	*outSlice = pagedItems
	*meta = api.PaginationMeta{
		TotalItems:  totalItems,
		TotalPages:  int(math.Ceil(float64(totalItems) / float64(req.Limit))),
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: endIdx < totalItems,
		HasPrevPage: req.Page > 1,
	}
}
