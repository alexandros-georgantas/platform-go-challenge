package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
