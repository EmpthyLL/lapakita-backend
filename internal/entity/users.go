package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PhoneNumberItem mewakili struktur satu item nomor telepon
type PhoneNumberItem struct {
	Number    string   `json:"number"`
	IsPrimary bool     `json:"is_primary"`
	Roles     []string `json:"roles"` // e.g. ["tenant", "stall_owner"]
}

// PhoneNumbers menampung daftar/array banyak nomor telepon per user
type PhoneNumbers []PhoneNumberItem

func (p PhoneNumbers) Value() (driver.Value, error) {
	if len(p) == 0 {
		return "[]", nil
	}
	return json.Marshal(p)
}

func (p *PhoneNumbers) Scan(value interface{}) error {
	if value == nil {
		*p = PhoneNumbers{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan PhoneNumbers: expected []byte or string, got %T", value)
	}
	return json.Unmarshal(bytes, p)
}

// GetPrimaryNumber mengembalikan nomor telepon utama user
func (p PhoneNumbers) GetPrimaryNumber() string {
	for _, num := range p {
		if num.IsPrimary {
			return num.Number
		}
	}
	if len(p) > 0 {
		return p[0].Number
	}
	return ""
}

// GetNumberForRole mengembalikan nomor spesifik untuk role/persona tertentu
func (p PhoneNumbers) GetNumberForRole(role string) string {
	for _, num := range p {
		for _, r := range num.Roles {
			if r == role {
				return num.Number
			}
		}
	}
	return p.GetPrimaryNumber()
}

// SetPrimaryNumber mengubah status primary ke nomor tertentu dan me-reset nomor lainnya
func (p *PhoneNumbers) SetPrimaryNumber(targetNumber string) {
	for i := range *p {
		if (*p)[i].Number == targetNumber {
			(*p)[i].IsPrimary = true
		} else {
			(*p)[i].IsPrimary = false
		}
	}
}

type RoleProfileItem struct {
	AvatarURL   string `json:"avatar_url"`
	DisplayName string `json:"display_name"`
}

type RoleProfiles map[string]RoleProfileItem

func (r RoleProfiles) Value() (driver.Value, error) {
	if len(r) == 0 {
		return "{}", nil
	}
	return json.Marshal(r)
}

func (r *RoleProfiles) Scan(value interface{}) error {
	if value == nil {
		*r = RoleProfiles{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("failed to scan RoleProfiles: expected []byte or string, got %T", value)
	}
	return json.Unmarshal(bytes, r)
}

type User struct {
	ID                    uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name                  string       `gorm:"type:varchar(255);not null" json:"name"`
	Email                 string       `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash          string       `gorm:"type:varchar(255);not null" json:"-"`
	DefaultAvatarURL      *string      `gorm:"type:text" json:"default_avatar_url"`
	PhoneNumbers          PhoneNumbers `gorm:"type:jsonb;not null;default:'[]'" json:"phone_numbers"`
	RoleProfiles          RoleProfiles `gorm:"type:jsonb;not null;default:'{}'" json:"role_profiles"`
	ActiveRole            string       `gorm:"type:varchar(32);default:'tenant'" json:"active_role"`
	SubscriptionPlan      string       `gorm:"type:varchar(32);default:'free'" json:"subscription_plan"`
	SubscriptionExpiresAt *time.Time   `gorm:"type:timestamp with time zone" json:"subscription_expires_at"`
	CreatedAt             time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time    `gorm:"autoUpdateTime" json:"updated_at"`

	IdentityProfile *UserIdentityProfile `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"identity_profile,omitempty"`
	BankAccounts    []BankAccount        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"bank_accounts,omitempty"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.ActiveRole == "" {
		u.ActiveRole = "tenant"
	}
	if u.SubscriptionPlan == "" {
		u.SubscriptionPlan = "free"
	}
	if u.PhoneNumbers == nil {
		u.PhoneNumbers = PhoneNumbers{}
	}
	if u.RoleProfiles == nil {
		u.RoleProfiles = RoleProfiles{}
	}
	return
}
