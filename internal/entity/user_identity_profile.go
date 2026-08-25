package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserIdentityProfile struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	FullNameKTP  string    `gorm:"type:varchar(255);not null" json:"full_name_ktp"`
	NIK          string    `gorm:"type:varchar(16);uniqueIndex;not null" json:"nik"`
	KTPPhotoURL  string    `gorm:"type:text;not null" json:"ktp_photo_url"`
	DomicileCity *string   `gorm:"type:varchar(128)" json:"domicile_city"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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
