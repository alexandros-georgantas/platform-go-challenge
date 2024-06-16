package models

import (
	"github.com/lib/pq"
)

type Chart struct {
	Base                Base `gorm:"embedded"`
	Title               string
	HorizontalAxisLabel string
	VerticalAxisLabel   string
	Data                pq.Int64Array `gorm:"type:integer[]"`
}
