package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/stall/dto"
	"lapakita-backend/pkg/api"
	"lapakita-backend/pkg/database"

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

func cleanLocationComponent(val string) string {
	cleaned := strings.TrimSpace(val)

	cleaned = strings.TrimPrefix(cleaned, "City of ")
	cleaned = strings.TrimPrefix(cleaned, "Regency of ")
	cleaned = strings.TrimPrefix(cleaned, "Province of ")
	cleaned = strings.TrimPrefix(cleaned, "Special Region of ")
	cleaned = strings.TrimPrefix(cleaned, "Special Capital Region of ")

	cleaned = strings.TrimPrefix(cleaned, "Kota ")
	cleaned = strings.TrimPrefix(cleaned, "Kabupaten ")
	cleaned = strings.TrimPrefix(cleaned, "Provinsi ")
	cleaned = strings.TrimPrefix(cleaned, "Propinsi ")
	cleaned = strings.TrimPrefix(cleaned, "Daerah Khusus Ibukota ")
	cleaned = strings.TrimPrefix(cleaned, "Daerah Istimewa ")

	cleaned = strings.TrimSuffix(cleaned, " City")
	cleaned = strings.TrimSuffix(cleaned, " Regency")

	return strings.TrimSpace(cleaned)
}

func translateDirection(val string) string {
	if val == "" {
		return val
	}

	directions := map[string]string{
		"North ":     " Utara",
		"South ":     " Selatan",
		"West ":      " Barat",
		"East ":      " Timur",
		"Central ":   " Tengah",
		"Southeast ": " Tenggara",
		"Southwest ": " Barat Daya",
		"Northeast ": " Timur Laut",
		"Northwest ": " Barat Daya",
	}

	for engPrefix, idSuffix := range directions {
		if strings.HasPrefix(val, engPrefix) {
			baseName := strings.TrimPrefix(val, engPrefix)
			return baseName + idSuffix
		}
	}

	return val
}

func normalizeLocationPart(value string) string {
	cleaned := cleanLocationComponent(value)
	translated := translateDirection(cleaned)
	return strings.ToLower(strings.TrimSpace(translated))
}

func applyLocationFilter(query *gorm.DB, location string) *gorm.DB {
	parts := strings.Split(location, ",")
	for _, rawPart := range parts {
		part := normalizeLocationPart(rawPart)
		if part == "" || part == "indonesia" {
			continue
		}

		likePart := "%" + part + "%"
		query = query.Where(
			"(LOWER(street_address) LIKE ? OR LOWER(suburb) LIKE ? OR LOWER(district) LIKE ? OR LOWER(city) LIKE ? OR LOWER(province) LIKE ? OR LOWER(country) LIKE ? OR LOWER(country_code) LIKE ?)",
			likePart, likePart, likePart, likePart, likePart, likePart, likePart,
		)
	}
	return query
}

func (r *StallRepository) Search(ctx context.Context, req dto.SearchStallRequest) ([]entity.Stall, api.PaginationMeta, error) {
	var stalls []entity.Stall
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).Model(&entity.Stall{}).Where("deleted_at IS NULL AND is_published = ?", true)
	query = query.Where("permanence_type != ? OR ((event_schedule->>'end_date')::date >= CURRENT_DATE)", entity.StallPermanenceTemporary)

	if req.Location != "" {
		query = applyLocationFilter(query, req.Location)
	}

	if req.PermanenceType != "" {
		query = query.Where("permanence_type = ?", req.PermanenceType)
	}

	if req.Placement != "" {
		query = query.Where("placement = ?", req.Placement)
	}

	if len(req.PropertyType) > 0 {
		query = query.Where("property_type IN ?", req.PropertyType)
	}

	if req.BusinessType != "" {
		query = query.Where("(permanence_type != 'temporary' OR (allowed_business_type_ids ::jsonb @> ? OR jsonb_array_length(allowed_business_type_ids) = 0))", fmt.Sprintf(`["%s"]`, req.BusinessType))
	}

	if req.PaymentCycle != "" {
		query = query.Where("allowed_payment_cycles ::jsonb @> ?", fmt.Sprintf(`["%s"]`, req.PaymentCycle))
	}

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

	if req.Capital != nil && *req.Capital > 0 {
		query = query.Where("(monthly_rate <= ? OR daily_rate <= ? OR yearly_rate <= ?)", *req.Capital, *req.Capital, *req.Capital)
	}
	if req.RentRange[1] > 0 {
		query = query.Where("monthly_rate BETWEEN ? AND ?", req.RentRange[0], req.RentRange[1])
	}

	if req.DepositRange[1] > 0 {
		query = query.Where("security_deposit BETWEEN ? AND ?", req.DepositRange[0], req.DepositRange[1])
	}

	if req.SizeRange[1] > 0 {
		query = query.Where("size_sqm BETWEEN ? AND ?", req.SizeRange[0], req.SizeRange[1])
	}
	if req.FloorCountRange[1] > 0 {
		query = query.Where("floor_level BETWEEN ? AND ?", req.FloorCountRange[0], req.FloorCountRange[1])
	}

	if req.OpeningTime != "" {
		query = query.Where("(operating_hours->>'opening_time') <= ?", req.OpeningTime)
	}
	if req.ClosingTime != "" {
		query = query.Where("(operating_hours->>'closing_time') >= ?", req.ClosingTime)
	}
	if req.Is24Hours != nil {
		query = query.Where("permanence_type = ? AND (operating_hours->>'is_24_hours')::boolean = ?", entity.StallPermanenceSemi, *req.Is24Hours)
	}

	if req.EventOperatingDays != "" {
		query = query.Where("event_operating_days = ?", req.EventOperatingDays)
	}
	if req.AttendanceRequirement != "" {
		query = query.Where("event_attendance_requirement = ?", req.AttendanceRequirement)
	}
	if req.CancellationPolicy != "" {
		query = query.Where("event_cancellation_policy = ?", req.CancellationPolicy)
	}
	if req.RegistrationDeadline != "" {
		query = query.Where("(event_schedule->>'registration_deadline')::date <= ?::date", req.RegistrationDeadline)
	}

	for _, facility := range req.Facilities {
		query = query.Where("facility_values ::jsonb @> ?", fmt.Sprintf(`["%s"]`, facility))
	}

	// Dynamic Order Clause
	var orderClause string
	switch req.SortBy {
	case "price-asc":
		orderClause = "COALESCE(monthly_rate, daily_rate, yearly_rate) ASC"
	case "price-desc":
		orderClause = "COALESCE(monthly_rate, daily_rate, yearly_rate) DESC"
	case "rating":
		query = query.Where("rating_avg >= 4.8")
		orderClause = "rating_avg DESC"
	case "reviews":
		orderClause = "review_count DESC"
	case "size-desc":
		orderClause = "size_sqm DESC"
	case "recommended":
		orderClause = "rating_avg DESC, review_count DESC"
	default:
		orderClause = "created_at DESC"
	}

	err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"stalls",
		&stalls,
		&meta,
		orderClause,
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return stalls, meta, nil
}

func (r *StallRepository) FindByOwnerID(ctx context.Context, req dto.GetOwnerStallsRequest) ([]entity.Stall, api.PaginationMeta, error) {
	var stalls []entity.Stall
	var meta api.PaginationMeta

	ownerUUID, err := uuid.Parse(req.OwnerID)
	if err != nil {
		return nil, api.PaginationMeta{}, err
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
		query = applyLocationFilter(query, req.Location)
	}
	if req.IsPublished != nil {
		query = query.Where("is_published = ?", *req.IsPublished)
	}

	err = database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"stalls",
		&stalls,
		&meta,
		"created_at DESC",
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return stalls, meta, nil
}

func (r *StallRepository) FindSimilar(ctx context.Context, currentStall *entity.Stall, req dto.GetSimilarStallsRequest) ([]entity.Stall, api.PaginationMeta, error) {
	var stalls []entity.Stall
	var meta api.PaginationMeta

	query := r.db.WithContext(ctx).Model(&entity.Stall{}).
		Where("deleted_at IS NULL AND is_published = ?", true).
		Where("id != ?", currentStall.ID)

	query = query.Where(
		"LOWER(city) = LOWER(?) OR property_type = ? OR permanence_type = ?",
		currentStall.City, currentStall.PropertyType, currentStall.PermanenceType,
	)

	orderByQuery := fmt.Sprintf(
		"CASE WHEN LOWER(city) = LOWER('%s') THEN 1 ELSE 2 END ASC, "+
			"CASE WHEN property_type = '%s' THEN 1 ELSE 2 END ASC, "+
			"rating_avg DESC, created_at DESC",
		strings.ReplaceAll(currentStall.City, "'", "''"),
		strings.ReplaceAll(currentStall.PropertyType, "'", "''"),
	)

	err := database.AutoPaginate(
		query,
		req.BasePaginationRequest,
		"stalls",
		&stalls,
		&meta,
		orderByQuery,
	)
	if err != nil {
		return nil, api.PaginationMeta{}, err
	}

	return stalls, meta, nil
}
