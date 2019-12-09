package models

import (
	"time"
)

// Base contains common columns for all tables.
type Base struct {
	ID        uint `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time `sql:"index"`
}
