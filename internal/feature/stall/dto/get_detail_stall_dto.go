package dto

type MultiPeriodPricing struct {
	DailyRate            *float64 `json:"dailyRate,omitempty"`
	MonthlyRate          *float64 `json:"monthlyRate,omitempty"`
	QuarterlyRate        *float64 `json:"quarterlyRate,omitempty"`
	SemesterlyRate       *float64 `json:"semesterlyRate,omitempty"`
	YearlyRate           *float64 `json:"yearlyRate,omitempty"`
	SecurityDeposit      float64  `json:"securityDeposit"`
	AllowedPaymentCycles []string `json:"allowedPaymentCycles"`
}

type NearbyLandmarkDTO struct {
	CategoryValue string  `json:"categoryValue"`
	Name          string  `json:"name"`
	DistanceKm    float64 `json:"distanceKm"`
}

type OwnerProfileSummary struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Contact     string  `json:"contact"`
	AvatarURL   string  `json:"avatarUrl"`
	Rating      float64 `json:"rating"`
	ReviewCount int     `json:"reviewCount"`
	JoinedYear  string  `json:"joinedYear"`
}

type FacilityImageDTO struct {
	ID      string `json:"id"`
	URL     string `json:"url"`
	Caption string `json:"caption"`
}

type StallMediaDTO struct {
	MainImage         string             `json:"mainImage"`
	FacilityImages    []FacilityImageDTO `json:"facilityImages"`
	VirtualTour360URL *string            `json:"virtualTour360Url,omitempty"`
}

type StallAddressDTO struct {
	Street         string `json:"street"`
	Suburb         string `json:"suburb"`
	District       string `json:"district"`
	City           string `json:"city"`
	Country        string `json:"country"`
	CountryCode    string `json:"countryCode"`
	Province       string `json:"province"`
	PostalCode     string `json:"postalCode"`
	MapURL         string `json:"mapUrl"`
	EmbeddedMapURL string `json:"embeddedMapUrl"`
}

type PermanentLeaseRules struct {
	MinimumLeaseMonths int      `json:"minimumLeaseMonths"`
	StartDateOptions   []string `json:"startDateOptions"`
	UtilityTerms       string   `json:"utilityTerms"`
}

type TemporaryLeaseRules struct {
	MinimumLeaseDays      int      `json:"minimumLeaseDays"`
	StartDateOptions      []string `json:"startDateOptions"`
	UtilityTerms          string   `json:"utilityTerms"`
	OperatingDays         string   `json:"operatingDays"`
	AttendanceRequirement string   `json:"attendanceRequirement"`
	CancellationPolicy    string   `json:"cancellationPolicy"`
}

type PermanentStallDetail struct {
	ID                    string              `json:"id"`
	Title                 string              `json:"title"`
	Description           string              `json:"description"`
	Media                 StallMediaDTO       `json:"media"`
	PropertyType          string              `json:"propertyType"`
	PropertyTypeValue     string              `json:"propertyTypeValue"`
	Placement             string              `json:"placement"`
	ElectricityCapacityVA int                 `json:"electricityCapacityVA"`
	Address               StallAddressDTO     `json:"address"`
	NearbyLandmarks       []NearbyLandmarkDTO `json:"nearbyLandmarks"`
	Pricing               MultiPeriodPricing  `json:"pricing"`
	FacilityValues        []string            `json:"facilityValues"`
	HouseRules            []string            `json:"houseRules"`
	Rating                float64             `json:"rating"`
	ReviewCount           int                 `json:"reviewCount"`
	Owner                 OwnerProfileSummary `json:"owner"`

	PermanenceType string  `json:"permanenceType"` // "permanent"
	SizeSqm        float64 `json:"sizeSqm"`
	Dimensions     struct {
		LengthMeters float64 `json:"lengthMeters"`
		WidthMeters  float64 `json:"widthMeters"`
	} `json:"dimensions"`
	FloorLevel int                 `json:"floorLevel"`
	LeaseRules PermanentLeaseRules `json:"leaseRules"`
}

type SemiPermanentStallDetail struct {
	ID                    string              `json:"id"`
	Title                 string              `json:"title"`
	Description           string              `json:"description"`
	Media                 StallMediaDTO       `json:"media"`
	PropertyType          string              `json:"propertyType"`
	PropertyTypeValue     string              `json:"propertyTypeValue"`
	Placement             string              `json:"placement"`
	ElectricityCapacityVA int                 `json:"electricityCapacityVA"`
	Address               StallAddressDTO     `json:"address"`
	NearbyLandmarks       []NearbyLandmarkDTO `json:"nearbyLandmarks"`
	Pricing               MultiPeriodPricing  `json:"pricing"`
	FacilityValues        []string            `json:"facilityValues"`
	HouseRules            []string            `json:"houseRules"`
	Rating                float64             `json:"rating"`
	ReviewCount           int                 `json:"reviewCount"`
	Owner                 OwnerProfileSummary `json:"owner"`

	PermanenceType    string `json:"permanenceType"` // "semi-permanent"
	ParentComplexName string `json:"parentComplexName"`
	OperatingHours    struct {
		OpeningTime string `json:"openingTime"`
		ClosingTime string `json:"closingTime"`
		Is24Hours   bool   `json:"is24Hours"`
	} `json:"operatingHours"`
	LeaseRules PermanentLeaseRules `json:"leaseRules"`
}

type TemporaryStallDetail struct {
	ID                    string              `json:"id"`
	Title                 string              `json:"title"`
	Description           string              `json:"description"`
	Media                 StallMediaDTO       `json:"media"`
	PropertyType          string              `json:"propertyType"`
	PropertyTypeValue     string              `json:"propertyTypeValue"`
	Placement             string              `json:"placement"`
	ElectricityCapacityVA int                 `json:"electricityCapacityVA"`
	Address               StallAddressDTO     `json:"address"`
	NearbyLandmarks       []NearbyLandmarkDTO `json:"nearbyLandmarks"`
	Pricing               MultiPeriodPricing  `json:"pricing"`
	FacilityValues        []string            `json:"facilityValues"`
	HouseRules            []string            `json:"houseRules"`
	Rating                float64             `json:"rating"`
	ReviewCount           int                 `json:"reviewCount"`
	Owner                 OwnerProfileSummary `json:"owner"`

	PermanenceType string `json:"permanenceType"` // "temporary"
	EventMeta      struct {
		EventName                      string `json:"eventName,omitempty"`
		EventStartDate                 string `json:"eventStartDate"`
		EventEndDate                   string `json:"eventEndDate"`
		RegistrationDeadlineDaysBefore int    `json:"registrationDeadlineDaysBefore"`
		TotalSlots                     int    `json:"totalSlots"`
		AvailableSlots                 int    `json:"availableSlots"`
	} `json:"eventMeta"`
	LeaseRules TemporaryLeaseRules `json:"leaseRules"`
}
