package models

// Drug defines a user for the app
type UsageHistory struct {
	Base
	DrugId   uint `gorm:"not null"`
	UserId   uint `gorm:"not null"`
	Quantity uint `gorm:"not null"`
}
