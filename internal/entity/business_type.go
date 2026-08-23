package entity

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessType struct {
	ID                         uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	LabelLang                  json.RawMessage `gorm:"type:jsonb;not null;default:'{\"en\": \"\", \"id\": \"\"}'" json:"label_lang"`
	GroupNameLang              json.RawMessage `gorm:"type:jsonb;not null;default:'{\"en\": \"\", \"id\": \"\"}'" json:"group_name_lang"`
	DefaultBEPMonths           int             `gorm:"type:int;not null;default:6" json:"default_bep_months"`
	DefaultCapital             float64         `gorm:"type:numeric(15,2);not null;default:35000000.00" json:"default_capital"`
	AvgGrossMarginRatio        float64         `gorm:"type:numeric(5,4);not null;default:0.5000" json:"avg_gross_margin_ratio"`
	IndustryRentToRevenueRatio float64         `gorm:"type:numeric(5,4);not null;default:0.1500" json:"industry_rent_to_revenue_ratio"`
	PermanencePresets          json.RawMessage `gorm:"type:jsonb;not null;default:'{}'" json:"permanence_presets"`
	RecommendedLandmarks       json.RawMessage `gorm:"type:jsonb;default:'[]'" json:"recommended_landmarks"`
	CreatedAt                  time.Time       `gorm:"autoUpdateTime" json:"created_at"`
	UpdatedAt                  time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                  gorm.DeletedAt  `gorm:"index" json:"deleted_at,omitempty"`
}

func (BusinessType) TableName() string {
	return "business_types"
}
