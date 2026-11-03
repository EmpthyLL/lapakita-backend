package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Business struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id" json:"id"`
	UserID         uuid.UUID      `gorm:"type:uuid;not null;index:idx_businesses_user_id" db:"user_id" json:"user_id"`
	BusinessTypeID uuid.UUID      `gorm:"type:uuid;not null" db:"business_type_id" json:"business_type_id"`
	Name           string         `gorm:"type:varchar(255);not null" db:"name" json:"name"`
	LogoURL        *string        `gorm:"type:text" db:"logo_url" json:"logo_url,omitempty"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index:idx_businesses_deleted_at" db:"deleted_at" json:"deleted_at,omitempty"`

	User         *User         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	BusinessType *BusinessType `gorm:"foreignKey:BusinessTypeID;constraint:OnDelete:RESTRICT" json:"business_type,omitempty"`
}

func (Business) TableName() string {
	return "businesses"
}

func (b *Business) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
