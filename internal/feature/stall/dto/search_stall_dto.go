package dto

import "lapakita-backend/pkg/api"

type SearchStallRequest struct {
	api.BasePaginationRequest
	Location                 string                `form:"location"`
	PermanenceType           string                `form:"permanenceType"`
	LandmarkEntries          []LandmarkRadiusEntry `form:"landmarkEntries"`
	PropertyType             []string              `form:"propertyType"`
	Placement                string                `form:"placement"`
	BusinessType             string                `form:"businessType"`
	Facilities               []string              `form:"facilities"`
	BepMonths                string                `form:"bepMonths"`
	CustomBepMonths          *float64              `form:"customBepMonths"`
	Capital                  *float64              `form:"capital"`
	RentRange                [2]float64            `form:"rentRange"`
	DepositRange             [2]float64            `form:"depositRange"`
	StartDate                string                `form:"startDate"`
	CustomStartDay           string                `form:"customStartDay"`
	MinLeasePeriod           string                `form:"minLeasePeriod"`
	CustomLeaseMonths        string                `form:"customLeaseMonths"`
	EventOperatingDays       string                `form:"eventOperatingDays"`
	AttendanceRequirement    string                `form:"attendanceRequirement"`
	CancellationPolicy       string                `form:"cancellationPolicy"`
	PaymentCycle             string                `form:"paymentCycle"`
	SizeRange                [2]float64            `form:"sizeRange"`
	FloorCountRange          [2]int                `form:"floorCountRange"`
	OpeningTime              string                `form:"openingTime"`
	ClosingTime              string                `form:"closingTime"`
	RegistrationDeadlineDays *int                  `form:"registrationDeadlineDays"`
	EventDurationDays        *int                  `form:"eventDurationDays"`
	SortBy                   string                `form:"sortBy"`
}

type StallResponseData struct {
	ID                      string             `json:"id"`
	StallOwnerID            string             `json:"stall_owner_id"`
	Title                   string             `json:"title"`
	Description             *string            `json:"description,omitempty"`
	PropertyType            string             `json:"property_type"`
	PermanenceType          string             `json:"permanence_type"`
	Placement               string             `json:"placement"`
	SizeSqm                 *float64           `json:"size_sqm,omitempty"`
	LengthMeters            *float64           `json:"length_meters,omitempty"`
	WidthMeters             *float64           `json:"width_meters,omitempty"`
	FloorLevel              *int               `json:"floor_level,omitempty"`
	ElectricityCapacityVA   *int               `json:"electricity_capacity_va,omitempty"`
	ParentComplexName       *string            `json:"parent_complex_name,omitempty"`
	OperatingHours          *OperatingHoursDTO `json:"operating_hours,omitempty"`
	EventSchedule           *EventScheduleDTO  `json:"event_schedule,omitempty"`
	SlotInfo                *SlotInfoDTO       `json:"slot_info,omitempty"`
	StreetAddress           string             `json:"street_address"`
	Suburb                  *string            `json:"suburb,omitempty"`
	District                *string            `json:"district,omitempty"`
	City                    string             `json:"city"`
	Province                string             `json:"province"`
	Country                 string             `json:"country"`
	CountryCode             string             `json:"country_code"`
	PostalCode              *string            `json:"postal_code,omitempty"`
	Latitude                *float64           `json:"latitude,omitempty"`
	Longitude               *float64           `json:"longitude,omitempty"`
	MapURL                  *string            `json:"map_url,omitempty"`
	EmbeddedMapURL          *string            `json:"embedded_map_url,omitempty"`
	NearbyLandmarks         []string           `json:"nearby_landmarks"`
	AllowedPaymentCycles    []string           `json:"allowed_payment_cycles"`
	DailyRate               *float64           `json:"daily_rate,omitempty"`
	MonthlyRate             *float64           `json:"monthly_rate,omitempty"`
	QuarterlyRate           *float64           `json:"quarterly_rate,omitempty"`
	SemesterlyRate          *float64           `json:"semesterly_rate,omitempty"`
	YearlyRate              *float64           `json:"yearly_rate,omitempty"`
	SecurityDeposit         float64            `json:"security_deposit"`
	MinimumLeaseMonths      *int               `json:"minimum_lease_months,omitempty"`
	MinimumLeaseDays        *int               `json:"minimum_lease_days,omitempty"`
	StartDateOptions        []string           `json:"start_date_options"`
	EventOperatingDays      *string            `json:"event_operating_days,omitempty"`
	EventAttendanceReq      *string            `json:"event_attendance_requirement,omitempty"`
	EventCancellationPolicy *string            `json:"event_cancellation_policy,omitempty"`
	UtilityTerms            *string            `json:"utility_terms,omitempty"`
	FacilityValues          []string           `json:"facility_values"`
	AllowedBusinessTypeIDs  []string           `json:"allowed_business_type_ids"`
	HouseRules              []string           `json:"house_rules"`
	DisplayMedia            DisplayMediaDTO    `json:"display_media"`
	LegalDocuments          []string           `json:"legal_documents"`
	RatingAvg               float64            `json:"rating_avg"`
	ReviewCount             int                `json:"review_count"`
	IsPublished             bool               `json:"is_published"`
	CreatedAt               string             `json:"created_at"`
	UpdatedAt               string             `json:"updated_at"`
}
