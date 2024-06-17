package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string
	GivenName    string
	Surname      string
	PasswordHash string
	Favorites    []Asset `gorm:"many2many:favorites;"`
}
