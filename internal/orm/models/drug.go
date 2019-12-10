package models

// Drug defines a user for the app
type Drug struct {
	Base
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}
