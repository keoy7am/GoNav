package models

import (
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Title       string `gorm:"not null"`
	URL         string `gorm:"not null"`
	Description string `gorm:"size:50"`
	IsPrivate   bool   `gorm:"default:false"`
	CategoryID  uint
	Category    Category
	UserID      uint
	Weight      int `gorm:"default:0"`
}
