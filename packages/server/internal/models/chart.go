package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Chart struct {
	gorm.Model
	Title               string
	HorizontalAxisLabel string
	VerticalAxisLabel   string
	Data                pq.Int64Array `gorm:"type:integer[]"`
	Asset               Asset         `gorm:"foreignKey:RelatedID"`
}
