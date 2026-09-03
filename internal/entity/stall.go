package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type StallPermanenceType string

const (
	StallPermanencePermanent StallPermanenceType = "permanent"
	StallPermanenceSemi      StallPermanenceType = "semi-permanent"
	StallPermanenceTemporary StallPermanenceType = "temporary"
)

type FacilityImage struct {
	URL string `json:"url"`
}

func (f *FacilityImage) UnmarshalJSON(data []byte) error {
	var url string
	if err := json.Unmarshal(data, &url); err == nil {
		f.URL = url
		return nil
	}

	type facilityImage FacilityImage
	var image facilityImage
	if err := json.Unmarshal(data, &image); err != nil {
		return err
	}
	*f = FacilityImage(image)
	return nil
}

type StallPlacementType string

const (
	StallPlacementIndoor  StallPlacementType = "indoor"
	StallPlacementOutdoor StallPlacementType = "outdoor"
	StallPlacementSemi    StallPlacementType = "semi-outdoor"
)

type EventOperatingDaysType string

const (
	EventOperatingDaysFullTime   EventOperatingDaysType = "full-time"
	EventOperatingDaysWeekends   EventOperatingDaysType = "weekends-only"
	EventOperatingDaysCustomDays EventOperatingDaysType = "custom-days"
)

type AttendanceRequirementType string

const (
	AttendanceRequirementMandatory AttendanceRequirementType = "mandatory"
	AttendanceRequirementOptional  AttendanceRequirementType = "optional"
)

type CancellationPolicyType string

const (
	CancellationPolicyStrict   CancellationPolicyType = "strict"
	CancellationPolicyFlexible CancellationPolicyType = "flexible"
	CancellationPolicyModerate CancellationPolicyType = "moderate"
)

type OperatingHours struct {
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
	Is24Hours   bool   `json:"is_24_hours"`
}

func (o OperatingHours) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *OperatingHours) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, o)
}

type EventSchedule struct {
	EventName                string `json:"event_name"`
	StartDate                string `json:"start_date"`
	EndDate                  string `json:"end_date"`
	RegistrationDeadlineDays int    `json:"registration_deadline_days"`
}

func (e EventSchedule) Value() (driver.Value, error) {
	return json.Marshal(e)
}

func (e *EventSchedule) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, e)
}

type SlotInfo struct {
	TotalSlots     int `json:"total_slots"`
	AvailableSlots int `json:"available_slots"`
}

func (s SlotInfo) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SlotInfo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, s)
}

type DisplayMedia struct {
	MainImage         string          `json:"mainImage"`
	FacilityImages    []FacilityImage `json:"facilityImages"`
	VirtualTour360URL *string         `json:"virtualTour360Url,omitempty"`
}

func (d DisplayMedia) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *DisplayMedia) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, d)
}

// NearbyLandmarkItem mewakili 1 item object di JSONB nearby_landmarks
type NearbyLandmarkItem struct {
	Name          string  `json:"name"`
	DistanceKm    float64 `json:"distanceKm"`
	CategoryValue string  `json:"categoryValue"`
}

type NearbyLandmarkArray []NearbyLandmarkItem

func (n NearbyLandmarkArray) Value() (driver.Value, error) {
	if n == nil {
		return json.Marshal([]NearbyLandmarkItem{})
	}
	return json.Marshal(n)
}

func (n *NearbyLandmarkArray) Scan(value interface{}) error {
	if value == nil {
		*n = NearbyLandmarkArray{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan NearbyLandmarkArray: %T", value)
	}
	return json.Unmarshal(bytes, n)
}

type StringArray []string

func (s StringArray) Value() (driver.Value, error) {
	if s == nil {
		return json.Marshal([]string{})
	}
	return json.Marshal(s)
}

func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan StringArray: %T", value)
	}
	return json.Unmarshal(bytes, s)
}

type Stall struct {
	ID                      uuid.UUID                  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id" json:"id"`
	StallOwnerID            uuid.UUID                  `gorm:"type:uuid;not null" db:"stall_owner_id" json:"stall_owner_id"`
	Title                   string                     `gorm:"type:varchar(255);not null" db:"title" json:"title"`
	Description             *string                    `gorm:"type:text" db:"description" json:"description"`
	PropertyType            string                     `gorm:"type:varchar(64);not null" db:"property_type" json:"property_type"`
	PermanenceType          StallPermanenceType        `gorm:"type:stall_permanence_type;not null;default:'permanent'" db:"permanence_type" json:"permanence_type"`
	Placement               StallPlacementType         `gorm:"type:stall_placement_type;not null;default:'indoor'" db:"placement" json:"placement"`
	SizeSqm                 *float64                   `gorm:"type:numeric(8,2)" db:"size_sqm" json:"size_sqm"`
	LengthMeters            *float64                   `gorm:"type:numeric(6,2)" db:"length_meters" json:"length_meters"`
	WidthMeters             *float64                   `gorm:"type:numeric(6,2)" db:"width_meters" json:"width_meters"`
	FloorLevel              *int                       `gorm:"type:int;default:1" db:"floor_level" json:"floor_level"`
	ElectricityCapacityVA   *int                       `gorm:"type:int;default:1300" db:"electricity_capacity_va" json:"electricity_capacity_va"`
	ParentComplexName       *string                    `gorm:"type:varchar(255)" db:"parent_complex_name" json:"parent_complex_name"`
	OperatingHours          *OperatingHours            `gorm:"type:jsonb" db:"operating_hours" json:"operating_hours"`
	EventSchedule           *EventSchedule             `gorm:"type:jsonb" db:"event_schedule" json:"event_schedule"`
	SlotInfo                *SlotInfo                  `gorm:"type:jsonb" db:"slot_info" json:"slot_info"`
	StreetAddress           string                     `gorm:"type:text;not null" db:"street_address" json:"street_address"`
	Suburb                  *string                    `gorm:"type:varchar(128)" db:"suburb" json:"suburb"`
	District                *string                    `gorm:"type:varchar(128)" db:"district" json:"district"`
	City                    string                     `gorm:"type:varchar(128);not null" db:"city" json:"city"`
	Province                string                     `gorm:"type:varchar(128);not null" db:"province" json:"province"`
	Country                 string                     `gorm:"type:varchar(128);not null;default:'Indonesia'" db:"country" json:"country"`
	CountryCode             string                     `gorm:"type:varchar(8);default:'ID'" db:"country_code" json:"country_code"`
	PostalCode              *string                    `gorm:"type:varchar(16)" db:"postal_code" json:"postal_code"`
	Latitude                *float64                   `gorm:"type:numeric(10,8)" db:"latitude" json:"latitude"`
	Longitude               *float64                   `gorm:"type:numeric(11,8)" db:"longitude" json:"longitude"`
	MapURL                  *string                    `gorm:"type:text" db:"map_url" json:"map_url"`
	EmbeddedMapURL          *string                    `gorm:"type:text" db:"embedded_map_url" json:"embedded_map_url"`
	NearbyLandmarks         NearbyLandmarkArray        `gorm:"type:jsonb;default:'[]'" db:"nearby_landmarks" json:"nearby_landmarks"`
	AllowedPaymentCycles    StringArray                `gorm:"type:jsonb;not null;default:'[\"month\"]'" db:"allowed_payment_cycles" json:"allowed_payment_cycles"`
	DailyRate               *float64                   `gorm:"type:numeric(15,2)" db:"daily_rate" json:"daily_rate"`
	MonthlyRate             *float64                   `gorm:"type:numeric(15,2)" db:"monthly_rate" json:"monthly_rate"`
	QuarterlyRate           *float64                   `gorm:"type:numeric(15,2)" db:"quarterly_rate" json:"quarterly_rate"`
	SemesterlyRate          *float64                   `gorm:"type:numeric(15,2)" db:"semesterly_rate" json:"semesterly_rate"`
	YearlyRate              *float64                   `gorm:"type:numeric(15,2)" db:"yearly_rate" json:"yearly_rate"`
	SecurityDeposit         float64                    `gorm:"type:numeric(15,2);not null;default:0.00" db:"security_deposit" json:"security_deposit"`
	MinimumLeaseMonths      *int                       `gorm:"type:int" db:"minimum_lease_months" json:"minimum_lease_months"`
	MinimumLeaseDays        *int                       `gorm:"type:int" db:"minimum_lease_days" json:"minimum_lease_days"`
	StartDateOptions        StringArray                `gorm:"type:jsonb;default:'[]'" db:"start_date_options" json:"start_date_options"`
	EventOperatingDays      *EventOperatingDaysType    `gorm:"type:event_operating_days_type" db:"event_operating_days" json:"event_operating_days"`
	EventAttendanceReq      *AttendanceRequirementType `gorm:"type:attendance_requirement_type" db:"event_attendance_requirement" json:"event_attendance_requirement"`
	EventCancellationPolicy *CancellationPolicyType    `gorm:"type:cancellation_policy_type" db:"event_cancellation_policy" json:"event_cancellation_policy"`
	UtilityTerms            *string                    `gorm:"type:text" db:"utility_terms" json:"utility_terms"`
	FacilityValues          StringArray                `gorm:"type:jsonb;default:'[]'" db:"facility_values" json:"facility_values"`
	AllowedBusinessTypeIDs  StringArray                `gorm:"type:jsonb;default:'[]'" db:"allowed_business_type_ids" json:"allowed_business_type_ids"`
	HouseRules              StringArray                `gorm:"type:jsonb;default:'[]'" db:"house_rules" json:"house_rules"`
	DisplayMedia            DisplayMedia               `gorm:"type:jsonb;not null" db:"display_media" json:"display_media"`
	LegalDocuments          StringArray                `gorm:"type:jsonb;default:'[]'" db:"legal_documents" json:"legal_documents"`
	RatingAvg               float64                    `gorm:"type:numeric(3,2);default:0.00" db:"rating_avg" json:"rating_avg"`
	ReviewCount             int                        `gorm:"type:int;default:0" db:"review_count" json:"review_count"`
	IsPublished             bool                       `gorm:"type:boolean;default:false" db:"is_published" json:"is_published"`
	FavoritedByUserIDs      StringArray                `gorm:"type:jsonb;default:'[]'" db:"favorited_by_user_ids" json:"favorited_by_user_ids"`
	CreatedAt               time.Time                  `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt               time.Time                  `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
	DeletedAt               *time.Time                 `gorm:"index" db:"deleted_at" json:"deleted_at,omitempty"`
}

func (Stall) TableName() string {
	return "stalls"
}
