package models

import (
	"gorm.io/gorm"
)

type UserSettings struct {
	gorm.Model
	UserID                uint   `gorm:"uniqueIndex"`
	EnablePublicBookmarks bool   `gorm:"default:false"`
	PublicID              string `gorm:"size:8;uniqueIndex"`
	EnableAccessCode      bool   `gorm:"default:false"`
	AccessCode            string
}
