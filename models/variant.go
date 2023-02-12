package models

import (
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID        uint   `gorm:"size:11;not null;uniqueIndex;primary_key" json:"-"`
	UniqueID  string `gorm:"size:250;not null;uniqueIndex" json:"-"`
	ProductID int    `gorm:"size:11;not null;" json:"-"`
	Product   Product
	Detail    []struct{}     `gorm:"not null" json:"detail"`
	Stock     int            `gorm:"size:11;not null" json:"stock"`
	Weight    int            `gorm:"size:11;not null" json:"weight"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
