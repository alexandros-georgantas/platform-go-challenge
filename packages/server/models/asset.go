package models

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Description string
	ObjectId    uuid.UUID `gorm:"type:uuid;"`
}
