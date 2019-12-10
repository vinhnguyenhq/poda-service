package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/vinhnguyenhq/poda-service/internal/orm/models"
	"gopkg.in/gormigrate.v1"
	"time"
)

func parseTime(layout string, dts string) time.Time {
	t, _ := time.Parse(layout, dts)
	return t
}

var (
	layout                                = "2006-01-02 15:04:05 -0700 MST"
	usageHistories []*models.UsageHistory = []*models.UsageHistory{
		&models.UsageHistory{UserId: 1, DrugId: 1, Quantity: 36, Base: models.Base{CreatedAt: parseTime(layout, "2019-11-25 12:00:00 +0000 UTC")}},
		&models.UsageHistory{UserId: 1, DrugId: 1, Quantity: 19, Base: models.Base{CreatedAt: parseTime(layout, "2019-11-26 12:00:00 +0000 UTC")}},
		&models.UsageHistory{UserId: 1, DrugId: 2, Quantity: 46, Base: models.Base{CreatedAt: parseTime(layout, "2019-11-26 12:00:00 +0000 UTC")}},
		&models.UsageHistory{UserId: 1, DrugId: 2, Quantity: 104, Base: models.Base{CreatedAt: parseTime(layout, "2019-11-27 12:00:00 +0000 UTC")}},
		&models.UsageHistory{UserId: 1, DrugId: 3, Quantity: 53, Base: models.Base{CreatedAt: parseTime(layout, "2019-11-27 12:00:00 +0000 UTC")}},
	}
)

// SeedUsageHistories
var SeedUsageHistories []*gormigrate.Migration = []*gormigrate.Migration{
	{
		ID: "seed_usage_history_1",
		Migrate: func(db *gorm.DB) error {
			return db.Create(usageHistories[0]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(usageHistories[0]).Error
		},
	},
	{
		ID: "seed_usage_history_2",
		Migrate: func(db *gorm.DB) error {
			return db.Create(usageHistories[1]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(usageHistories[1]).Error
		},
	},
	{
		ID: "seed_usage_history_3",
		Migrate: func(db *gorm.DB) error {
			return db.Create(usageHistories[2]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(usageHistories[2]).Error
		},
	},
	{
		ID: "seed_usage_history_4",
		Migrate: func(db *gorm.DB) error {
			return db.Create(usageHistories[3]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(usageHistories[3]).Error
		},
	},
	{
		ID: "seed_usage_history_5",
		Migrate: func(db *gorm.DB) error {
			return db.Create(usageHistories[4]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(usageHistories[4]).Error
		},
	},
}
