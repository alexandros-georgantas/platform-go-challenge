package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Email        string
	GivenName    string
	Surname      string
	PasswordHash string
}
