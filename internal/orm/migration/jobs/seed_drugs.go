package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/vinhnguyenhq/poda-service/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	drugs []*models.Drug = []*models.Drug{
		&models.Drug{Base: models.Base{ID: 1}, Name: "Buprenorphine", Description: "BuTrans"},
		&models.Drug{Base: models.Base{ID: 2}, Name: "Buprenorphine-naloxone", Description: "Suboxone"},
		&models.Drug{Base: models.Base{ID: 3}, Name: "Codeine", Description: "Tylenol 2,3,4"},
		&models.Drug{Base: models.Base{ID: 4}, Name: "Fentanyl", Description: "Abstral, Duragesic, Onsolis"},
		&models.Drug{Base: models.Base{ID: 5}, Name: "Hydrocodone", Description: "Tussionex, Vicoprofen"},
		&models.Drug{Base: models.Base{ID: 6}, Name: "Hydromorphone", Description: "Dilaudid"},
		&models.Drug{Base: models.Base{ID: 7}, Name: "Meperidine", Description: "Demerol"},
		&models.Drug{Base: models.Base{ID: 8}, Name: "Methadone", Description: "Methadose, Metadol"},
		&models.Drug{Base: models.Base{ID: 9}, Name: "Morphine", Description: "Doloral, Statex, M.O.S."},
		&models.Drug{Base: models.Base{ID: 10}, Name: "Oxycodone", Description: "OxyNEO, Percocet, Oxycocet Percodan"},
		&models.Drug{Base: models.Base{ID: 11}, Name: "Pentazocine", Description: "Talwin"},
		&models.Drug{Base: models.Base{ID: 12}, Name: "Tapentadol", Description: "Nucynta"},
		&models.Drug{Base: models.Base{ID: 13}, Name: "Tramadol", Description: "Ultram Tramacet Tridural Durela"},
	}
)

// SeedDrugs
var SeedDrugs []*gormigrate.Migration = []*gormigrate.Migration{
	{
		ID: "seed_drug_1",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[0]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[0]).Error
		},
	},
	{
		ID: "seed_drug_2",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[1]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[1]).Error
		},
	},
	{
		ID: "seed_drug_3",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[2]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[2]).Error
		},
	},
	{
		ID: "seed_drug_4",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[3]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[3]).Error
		},
	},
	{
		ID: "seed_drug_5",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[4]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[4]).Error
		},
	},
	{
		ID: "seed_drug_6",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[5]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[5]).Error
		},
	},
	{
		ID: "seed_drug_7",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[6]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[6]).Error
		},
	},
	{
		ID: "seed_drug_8",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[7]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[7]).Error
		},
	},
	{
		ID: "seed_drug_9",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[8]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[8]).Error
		},
	},
	{
		ID: "seed_drug_10",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[9]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[9]).Error
		},
	},
	{
		ID: "seed_drug_11",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[10]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[10]).Error
		},
	},
	{
		ID: "seed_drug_12",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[11]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[11]).Error
		},
	},
	{
		ID: "seed_drug_13",
		Migrate: func(db *gorm.DB) error {
			return db.Create(drugs[12]).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Delete(drugs[12]).Error
		},
	},
}
