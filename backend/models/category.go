package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Icon      string
	IsPrivate bool `gorm:"default:false"`
	Weight    int  `gorm:"default:0"`
	UserID    uint
	Bookmarks []Bookmark
}
