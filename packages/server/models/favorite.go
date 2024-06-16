package models

import "github.com/google/uuid"

type Favorite struct {
	Base    Base      `gorm:"embedded"`
	UserId  uuid.UUID `gorm:"type:uuid;uniqueIndex:user_asset,priority:1"`
	AssetId uuid.UUID `gorm:"type:uuid;uniqueIndex:user_asset,priority:2"`
}
