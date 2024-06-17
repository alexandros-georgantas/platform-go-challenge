package models

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"primaryKey;uniqueIndex:user_asset,priority:1"`
	AssetID uint `gorm:"primaryKey;uniqueIndex:user_asset,priority:2"`
}
