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

type StallLocationSummary struct {
	Area        string `json:"area"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
}

type BaseStallResponse struct {
	ID                     string               `json:"id"`
	Title                  string               `json:"title"`
	ImageURL               string               `json:"imageUrl"`
	Location               StallLocationSummary `json:"location"`
	PropertyType           string               `json:"propertyType"`
	Placement              string               `json:"placement"`
	CheapestPriceFormatted string               `json:"cheapestPriceFormatted"`
	CheapestPricePeriod    string               `json:"cheapestPricePeriod"`
	Rating                 float64              `json:"rating"`
	ReviewCount            int                  `json:"reviewCount"`
}

type PermanentStallResponse struct {
	BaseStallResponse
	PermanenceType string `json:"permanenceType"` // "permanent"
	Space          struct {
		SizeSqm    float64 `json:"sizeSqm"`
		FloorCount int     `json:"floorCount"`
	} `json:"space"`
}

type SemiPermanentStallResponse struct {
	BaseStallResponse
	PermanenceType string `json:"permanenceType"` // "semi-permanent"
	OperatingHours struct {
		Open  string `json:"open"`
		Close string `json:"close"`
	} `json:"operatingHours"`
}

type TemporaryStallResponse struct {
	BaseStallResponse
	PermanenceType string `json:"permanenceType"` // "temporary"
	Event          struct {
		RegistrationDeadlineDays int `json:"registrationDeadlineDays"`
		DurationDays             int `json:"durationDays"`
	} `json:"event"`
}
