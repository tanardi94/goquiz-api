package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"size:11;not null;uniqueIndex;primary_key" json:"-"`
	UniqueID    string `gorm:"size:250;not null;uniqueIndex" json:"-"`
	CategoryID  int    `gorm:"size:11;not null;" json:"category_id"`
	Category    Category
	Name        string         `gorm:"size:250;not null" json:"name"`
	Description string         `gorm:"not null;" json:"description"`
	Code        string         `gorm:"size:20" json:"code"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}
