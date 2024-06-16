package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Chart struct {
	ID                  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Title               string
	HorizontalAxisLabel string
	VerticalAxisLabel   string
	Data                pq.Int64Array `gorm:"type:integer[]"`
}
