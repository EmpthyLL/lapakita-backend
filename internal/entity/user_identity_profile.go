package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserIdentityProfile struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID           uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	DocumentType     string    `gorm:"type:varchar(32);not null;default:'national_id'" json:"document_type"`
	FullNameIdentity string    `gorm:"type:varchar(255);not null" json:"full_name_identity"`
	DocumentNumber   string    `gorm:"type:varchar(64);not null" json:"document_number"`
	DocumentPhotoURL string    `gorm:"type:text;not null" json:"document_photo_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (UserIdentityProfile) TableName() string {
	return "user_identity_profiles"
}

func (u *UserIdentityProfile) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
