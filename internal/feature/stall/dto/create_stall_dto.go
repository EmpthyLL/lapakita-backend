package dto

type CreateStallRequest struct {
	Title                   string             `json:"title" binding:"required"`
	Description             *string            `json:"description"`
	PropertyType            string             `json:"property_type" binding:"required"`
	PermanenceType          string             `json:"permanence_type" binding:"required"`
	Placement               string             `json:"placement" binding:"required"`
	SizeSqm                 *float64           `json:"size_sqm"`
	LengthMeters            *float64           `json:"length_meters"`
	WidthMeters             *float64           `json:"width_meters"`
	FloorLevel              *int               `json:"floor_level"`
	ElectricityCapacityVA   *int               `json:"electricity_capacity_va"`
	ParentComplexName       *string            `json:"parent_complex_name"`
	OperatingHours          *OperatingHoursDTO `json:"operating_hours"`
	EventSchedule           *EventScheduleDTO  `json:"event_schedule"`
	SlotInfo                *SlotInfoDTO       `json:"slot_info"`
	StreetAddress           string             `json:"street_address" binding:"required"`
	Suburb                  *string            `json:"suburb"`
	District                *string            `json:"district"`
	City                    string             `json:"city" binding:"required"`
	Province                string             `json:"province" binding:"required"`
	Country                 string             `json:"country"`
	CountryCode             string             `json:"country_code"`
	PostalCode              *string            `json:"postal_code"`
	Latitude                *float64           `json:"latitude"`
	Longitude               *float64           `json:"longitude"`
	MapURL                  *string            `json:"map_url"`
	EmbeddedMapURL          *string            `json:"embedded_map_url"`
	NearbyLandmarks         []string           `json:"nearby_landmarks"`
	AllowedPaymentCycles    []string           `json:"allowed_payment_cycles" binding:"required"`
	DailyRate               *float64           `json:"daily_rate"`
	MonthlyRate             *float64           `json:"monthly_rate"`
	QuarterlyRate           *float64           `json:"quarterly_rate"`
	SemesterlyRate          *float64           `json:"semesterly_rate"`
	YearlyRate              *float64           `json:"yearly_rate"`
	SecurityDeposit         float64            `json:"security_deposit"`
	MinimumLeaseMonths      *int               `json:"minimum_lease_months"`
	MinimumLeaseDays        *int               `json:"minimum_lease_days"`
	StartDateOptions        []string           `json:"start_date_options"`
	EventOperatingDays      *string            `json:"event_operating_days"`
	EventAttendanceReq      *string            `json:"event_attendance_requirement"`
	EventCancellationPolicy *string            `json:"event_cancellation_policy"`
	UtilityTerms            *string            `json:"utility_terms"`
	FacilityValues          []string           `json:"facility_values"`
	AllowedBusinessTypeIDs  []string           `json:"allowed_business_type_ids"`
	HouseRules              []string           `json:"house_rules"`
	DisplayMedia            DisplayMediaDTO    `json:"display_media" binding:"required"`
	LegalDocuments          []string           `json:"legal_documents"`
	IsPublished             bool               `json:"is_published"`
}
