package models

// User defines a user for the app
type User struct {
	Base
	Alias string `gorm:"not null;unique_index:idx_email"`
}
