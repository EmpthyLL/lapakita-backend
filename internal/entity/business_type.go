package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// I18nText menampung teks multi-bahasa (e.g. English & Indonesian)
type I18nText struct {
	EN string `json:"en"`
	ID string `json:"id"`
}

func (i I18nText) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *I18nText) Scan(value interface{}) error {
	if value == nil {
		*i = I18nText{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan I18nText: %T", value)
	}
	return json.Unmarshal(bytes, i)
}

// RecommendedLandmarkItem mewakili 1 rekomendasi landmark dalam JSONB business_types
type RecommendedLandmarkItem struct {
	CategoryValue string `json:"category_value"`
	Name          string `json:"name"`
}

type RecommendedLandmarkArray []RecommendedLandmarkItem

func (r RecommendedLandmarkArray) Value() (driver.Value, error) {
	if r == nil {
		return json.Marshal([]RecommendedLandmarkItem{})
	}
	return json.Marshal(r)
}

func (r *RecommendedLandmarkArray) Scan(value interface{}) error {
	if value == nil {
		*r = RecommendedLandmarkArray{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan RecommendedLandmarkArray: %T", value)
	}
	return json.Unmarshal(bytes, r)
}

type BusinessType struct {
	ID                         uuid.UUID                `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id" json:"id"`
	LabelLang                  I18nText                 `gorm:"type:jsonb;not null;default:'{\"en\": \"\", \"id\": \"\"}'" db:"label_lang" json:"label_lang"`
	GroupNameLang              I18nText                 `gorm:"type:jsonb;not null;default:'{\"en\": \"\", \"id\": \"\"}'" db:"group_name_lang" json:"group_name_lang"`
	DefaultBEPMonths           int                      `gorm:"type:int;not null;default:6" db:"default_bep_months" json:"default_bep_months"`
	DefaultCapital             float64                  `gorm:"type:numeric(15,2);not null;default:35000000.00" db:"default_capital" json:"default_capital"`
	AvgGrossMarginRatio        float64                  `gorm:"type:numeric(5,4);not null;default:0.5000" db:"avg_gross_margin_ratio" json:"avg_gross_margin_ratio"`
	IndustryRentToRevenueRatio float64                  `gorm:"type:numeric(5,4);not null;default:0.1500" db:"industry_rent_to_revenue_ratio" json:"industry_rent_to_revenue_ratio"`
	PermanencePresets          json.RawMessage          `gorm:"type:jsonb;not null;default:'{}'" db:"permanence_presets" json:"permanence_presets"`
	RecommendedLandmarks       RecommendedLandmarkArray `gorm:"type:jsonb;default:'[]'" db:"recommended_landmarks" json:"recommended_landmarks"`
	CreatedAt                  time.Time                `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt                  time.Time                `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
	DeletedAt                  gorm.DeletedAt           `gorm:"index" db:"deleted_at" json:"deleted_at,omitempty"`
}

func (BusinessType) TableName() string {
	return "business_types"
}

func (b *BusinessType) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
