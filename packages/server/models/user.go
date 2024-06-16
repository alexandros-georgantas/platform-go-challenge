package models

type User struct {
	Base         Base `gorm:"embedded"`
	Email        string
	GivenName    string
	Surname      string
	PasswordHash string
}
