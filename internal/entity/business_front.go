package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessFront struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id" json:"id"`
	BusinessID uuid.UUID      `gorm:"type:uuid;not null;index:idx_business_fronts_business_id" db:"business_id" json:"business_id"`
	Platform   string         `gorm:"type:varchar(64);not null" db:"platform" json:"platform"`
	StoreName  string         `gorm:"type:varchar(255);not null" db:"store_name" json:"store_name"`
	StoreURL   *string        `gorm:"type:text" db:"store_url" json:"store_url,omitempty"`
	IsActive   bool           `gorm:"type:boolean;default:true" db:"is_active" json:"is_active"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index:idx_business_fronts_deleted_at" db:"deleted_at" json:"deleted_at,omitempty"`

	Business *Business `gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE" json:"business,omitempty"`
}

func (BusinessFront) TableName() string {
	return "business_fronts"
}

func (bf *BusinessFront) BeforeCreate(tx *gorm.DB) (err error) {
	if bf.ID == uuid.Nil {
		bf.ID = uuid.New()
	}
	return
}
