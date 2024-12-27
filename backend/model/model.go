package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type User struct {
	ID                 uuid.UUID      `gorm:"type:uuid" json:"id"`
	Aud                string         `json:"aud"`
	Role               string         `json:"role"`
	Email              string         `json:"email"`
	InvitedAt          time.Time      `json:"invited_at"`
	ConfirmedAt        time.Time      `json:"-"`
	EmailConfirmedAt   time.Time      `json:"email_confirmed_at,omitempty"`
	ConfirmationSentAt time.Time      `json:"confirmation_sent_at,omitempty"`
	AppMetadata        datatypes.JSON `json:"app_metadata"`
	UserMetadata       datatypes.JSON `json:"user_metadata"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	Phone              string         `gorm:"unique" json:"phone"`
	Blogs              []Blog         `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"blogs"`
}

func (u *User) TableName() string {
    return "auth.users"
}

type Blog struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
