package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"size:11;not null;uniqueIndex;primary_key" json:"-"`
	UniqueID  string         `gorm:"size:250;not null;uniqueIndex" json:"unique_id"`
	Name      string         `gorm:"size:250;not null" json:"name"`
	ParentID  int            `gorm:"size:20" json:"parent"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
