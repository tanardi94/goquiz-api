package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"size:11;not null;uniqueIndex;primary_key" json:"-"`
	UniqueID        string         `gorm:"size:250;not null;uniqueIndex" json:"-"`
	Username        string         `gorm:"size:250;not null;uniqueIndex" json:"username"`
	Name            string         `gorm:"size:250;not null" json:"name"`
	Email           string         `gorm:"size:250;not null;uniqueIndex" json:"email"`
	Password        string         `gorm:"size:250;not null" json:"-"`
	Phone           string         `gorm:"size:20" json:"phone"`
	EmailVerifiedAt time.Time      `json:"email_verified_at"`
	PhoneVerifiedAt time.Time      `json:"phone_verified_at"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}
