package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/stall/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StallRepository struct {
	db *gorm.DB
}

func NewStallRepository(db *gorm.DB) *StallRepository {
	return &StallRepository{db: db}
}

func (r *StallRepository) Create(ctx context.Context, stall *entity.Stall) error {
	return r.db.WithContext(ctx).Create(stall).Error
}

func (r *StallRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Stall, error) {
	var stall entity.Stall
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").First(&stall, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &stall, err
}

func (r *StallRepository) Update(ctx context.Context, stall *entity.Stall) error {
	return r.db.WithContext(ctx).Save(stall).Error
}

func (r *StallRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&entity.Stall{}).Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *StallRepository) Search(ctx context.Context, req dto.SearchStallRequest) ([]entity.Stall, int64, error) {
	var stalls []entity.Stall
	var total int64

	query := r.db.WithContext(ctx).Model(&entity.Stall{}).Where("deleted_at IS NULL AND is_published = ?", true)

	// 1. Filter Combined Location
	if req.Location != "" {
		loc := "%" + strings.ToLower(req.Location) + "%"
		query = query.Where("(LOWER(street_address) LIKE ? OR LOWER(suburb) LIKE ? OR LOWER(district) LIKE ? OR LOWER(city) LIKE ? OR LOWER(province) LIKE ?)", loc, loc, loc, loc, loc)
	}

	// 2. Filter Permanence Type
	if req.PermanenceType != "" {
		query = query.Where("permanence_type = ?", req.PermanenceType)
	}

	// 3. Filter Placement
	if req.Placement != "" {
		query = query.Where("placement = ?", req.Placement)
	}

	// 4. Filter Property Types
	if len(req.PropertyType) > 0 {
		query = query.Where("property_type IN ?", req.PropertyType)
	}

	// 5. Filter Business Type
	if req.BusinessType != "" {
		query = query.Where("allowed_business_type_ids ::jsonb @> ?", fmt.Sprintf(`["%s"]`, req.BusinessType))
	}

	// 6. Filter Payment Cycle
	if req.PaymentCycle != "" {
		query = query.Where("allowed_payment_cycles ::jsonb @> ?", fmt.Sprintf(`["%s"]`, req.PaymentCycle))
	}

	// 7. Filter Nearby Landmark Entries (JSONB Array Search)
	for _, lm := range req.LandmarkEntries {
		if lm.Landmark != "" {
			lmSearch := "%" + strings.ToLower(lm.Landmark) + "%"
			query = query.Where("EXISTS (SELECT 1 FROM jsonb_array_elements(nearby_landmarks) elem WHERE LOWER(elem->>'name') LIKE ?)", lmSearch)
		}
		if lm.Radius != "" {
			if radiusFloat, err := strconv.ParseFloat(lm.Radius, 64); err == nil && radiusFloat > 0 {
				query = query.Where("EXISTS (SELECT 1 FROM jsonb_array_elements(nearby_landmarks) elem WHERE (elem->>'distanceKm')::numeric <= ?)", radiusFloat)
			}
		}
	}

	// 8. Filter Capital & Rent Range
	if req.Capital != nil && *req.Capital > 0 {
		query = query.Where("(monthly_rate <= ? OR daily_rate <= ? OR yearly_rate <= ?)", *req.Capital, *req.Capital, *req.Capital)
	}
	if req.RentRange[1] > 0 {
		query = query.Where("monthly_rate BETWEEN ? AND ?", req.RentRange[0], req.RentRange[1])
	}

	// 9. Filter Security Deposit Range
	if req.DepositRange[1] > 0 {
		query = query.Where("security_deposit BETWEEN ? AND ?", req.DepositRange[0], req.DepositRange[1])
	}

	// 10. Filter Size Range & Floor Count Range
	if req.SizeRange[1] > 0 {
		query = query.Where("size_sqm BETWEEN ? AND ?", req.SizeRange[0], req.SizeRange[1])
	}
	if req.FloorCountRange[1] > 0 {
		query = query.Where("floor_level BETWEEN ? AND ?", req.FloorCountRange[0], req.FloorCountRange[1])
	}

	// 11. Filter Operating Hours
	if req.OpeningTime != "" {
		query = query.Where("(operating_hours->>'opening_time') <= ?", req.OpeningTime)
	}
	if req.ClosingTime != "" {
		query = query.Where("(operating_hours->>'closing_time') >= ?", req.ClosingTime)
	}

	// 12. Filter Event Specific Terms
	if req.EventOperatingDays != "" {
		query = query.Where("event_operating_days = ?", req.EventOperatingDays)
	}
	if req.AttendanceRequirement != "" {
		query = query.Where("event_attendance_requirement = ?", req.AttendanceRequirement)
	}
	if req.CancellationPolicy != "" {
		query = query.Where("event_cancellation_policy = ?", req.CancellationPolicy)
	}
	if req.RegistrationDeadlineDays != nil {
		query = query.Where("(event_schedule->>'registration_deadline_days')::int <= ?", *req.RegistrationDeadlineDays)
	}

	// 13. Filter Facilities
	for _, facility := range req.Facilities {
		query = query.Where("facility_values ::jsonb @> ?", fmt.Sprintf(`["%s"]`, facility))
	}

	// Count Total Matching Stalls
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination Setup
	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	// 14. Sorting Options
	switch req.SortBy {
	case "price-asc":
		query = query.Order("COALESCE(monthly_rate, daily_rate, yearly_rate) ASC")
	case "price-desc":
		query = query.Order("COALESCE(monthly_rate, daily_rate, yearly_rate) DESC")
	case "rating":
		query = query.Where("rating_avg >= 4.8").Order("rating_avg DESC")
	case "reviews":
		query = query.Order("review_count DESC")
	case "size-desc":
		query = query.Order("size_sqm DESC")
	case "recommended":
		query = query.Order("rating_avg DESC, review_count DESC")
	default:
		query = query.Order("created_at DESC")
	}

	err := query.Offset(offset).Limit(limit).Find(&stalls).Error
	return stalls, total, err
}

func (r *StallRepository) FindByOwnerID(ctx context.Context, req dto.GetOwnerStallsRequest) ([]entity.Stall, int64, error) {
	var stalls []entity.Stall
	var total int64

	ownerUUID, err := uuid.Parse(req.OwnerID)
	if err != nil {
		return nil, 0, err
	}

	query := r.db.WithContext(ctx).Model(&entity.Stall{}).Where("deleted_at IS NULL AND stall_owner_id = ?", ownerUUID)

	if req.Title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(req.Title)+"%")
	}
	if req.PropertyType != "" {
		query = query.Where("property_type = ?", req.PropertyType)
	}
	if req.PermanenceType != "" {
		query = query.Where("permanence_type = ?", req.PermanenceType)
	}
	if req.Placement != "" {
		query = query.Where("placement = ?", req.Placement)
	}
	if req.Location != "" {
		loc := "%" + strings.ToLower(req.Location) + "%"
		query = query.Where("(LOWER(street_address) LIKE ? OR LOWER(suburb) LIKE ? OR LOWER(district) LIKE ? OR LOWER(city) LIKE ? OR LOWER(province) LIKE ?)", loc, loc, loc, loc, loc)
	}
	if req.IsPublished != nil {
		query = query.Where("is_published = ?", *req.IsPublished)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	limit := req.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	err = query.Order("created_at DESC").Offset(offset).Limit(limit).Find(&stalls).Error
	return stalls, total, err
}
