package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/vinhnguyenhq/poda-service/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	users []*models.User = []*models.User{
		&models.User{Base: models.Base{ID: 1}, Alias: "rick"},
		&models.User{Base: models.Base{ID: 2}, Alias: "daryl"},
	}
)

// SeedUsers
var SeedUsers []*gormigrate.Migration = []*gormigrate.Migration{
	{
		ID: "seed_user_1",
		Migrate: func(db *gorm.DB) error {
			return db.Create(users[0]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(users[0]).Error
		},
	},
	{
		ID: "seed_user_2",
		Migrate: func(db *gorm.DB) error {
			return db.Create(users[1]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(users[1]).Error
		},
	},
}
