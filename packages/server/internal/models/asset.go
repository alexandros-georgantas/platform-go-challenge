package models

import (
	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Description string
	RelatedType string `gorm:"index"` // Polymorphic type field
	RelatedID   uint   `gorm:"index"` // Polymorphic ID field
	// RelatedChart    *Chart    `gorm:"polymorphic:Related;polymorphicValue:chart"`
	// RelatedInsight  *Insight  `gorm:"polymorphic:Related;polymorphicValue:insight"`
	// RelatedAudience *Audience `gorm:"polymorphic:Related;polymorphicValue:audience"`
	Users []User `gorm:"many2many:favorites;"`
}
