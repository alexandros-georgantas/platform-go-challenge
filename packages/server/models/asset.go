package models

import (
	"github.com/google/uuid"
)

type Asset struct {
	Base        Base `gorm:"embedded"`
	Description string
	ObjectId    uuid.UUID `gorm:"type:uuid"`
}
