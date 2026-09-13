package database

import (
	"fmt"
	"lapakita-backend/pkg/api"

	"gorm.io/gorm"
)

// AutoPaginate memproses query secara otomatis berdasarkan isi BasePaginationRequest.
// customOrderOpsional opsional jika query membutuhkan urutan khusus (e.g. JSON extraction order).
// primaryKeyColumn opsional, jika kosong secara default akan menggunakan "id".
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

	orderClause := "created_at DESC"
	if customOrderOpsional != "" {
		orderClause = customOrderOpsional
	}

	var totalItems int64
	countQuery := db.Session(&gorm.Session{})
	if err := countQuery.Count(&totalItems).Error; err != nil {
		return err
	}

	query := db.Session(&gorm.Session{})

	// -------------------------------------------------------------------------
	// KASUS A: Membawa Selected ID (Dropdown Hydration / Anchor Pagination)
	// -------------------------------------------------------------------------
	if req.SelectedID != "" {
		var anchor T
		if err := db.Where(fmt.Sprintf("%s = ?", pk), req.SelectedID).First(&anchor).Error; err == nil {
			if req.Page > 1 {
				offset := (req.Page - 1) * req.Limit
				query = query.Where(fmt.Sprintf("%s <= (SELECT %s FROM %s WHERE %s = ?)", pk, pk, targetTableName, pk), req.SelectedID).
					Offset(offset)
			} else if req.Page < 1 {
				pageOffset := (-req.Page) * req.Limit
				query = query.Where(fmt.Sprintf("%s > (SELECT %s FROM %s WHERE %s = ?)", pk, pk, targetTableName, pk), req.SelectedID).
					Offset(pageOffset - req.Limit)
			} else {
				query = query.Where(fmt.Sprintf("%s <= (SELECT %s FROM %s WHERE %s = ?)", pk, pk, targetTableName, pk), req.SelectedID)
			}
		}
	} else {
		// -------------------------------------------------------------------------
		// KASUS B: Paginate Biasa (Offset Base standard)
		// -------------------------------------------------------------------------
		offset := (req.Page - 1) * req.Limit
		query = query.Offset(offset)
	}

	err := query.Limit(req.Limit + 1).Order(orderClause).Find(outSlice).Error
	if err != nil {
		return err
	}

	items := *outSlice
	hasNext := false
	if len(items) > req.Limit {
		hasNext = true
		*outSlice = items[:req.Limit]
	}

	totalPages := int((totalItems + int64(req.Limit) - 1) / int64(req.Limit))

	*meta = api.PaginationMeta{
		TotalItems:  int(totalItems),
		TotalPages:  totalPages,
		CurrentPage: req.Page,
		PerPage:     req.Limit,
		HasNextPage: hasNext,
		HasPrevPage: req.Page > 1 || req.Page < 0,
	}

	return nil
}
