package resolvers

import (
	"context"

	"github.com/vinhnguyenhq/poda-service/internal/gql/models"
	tf "github.com/vinhnguyenhq/poda-service/internal/gql/resolvers/transformations"
	log "github.com/vinhnguyenhq/poda-service/internal/logger"
	dbm "github.com/vinhnguyenhq/poda-service/internal/orm/models"
)

// Drugs
func (r *queryResolver) Drugs(ctx context.Context) ([]*models.Drug, error) {
	return drugs(r)
}

// Drug
func (r *queryResolver) Drug(ctx context.Context, id string) (*models.Drug, error) {
	return drug(r, id)
}

// ## Helper functions
func drugs(r *queryResolver) ([]*models.Drug, error) {
	entity := "drugs"
	record := []*models.Drug{}
	dbRecords := []*dbm.Drug{}

	db := r.ORM.DB.New()
	db = db.Find(&dbRecords)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBDrugToGQLDrug(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record = append(record, rec)
		}
	}

	return record, db.Error
}

func drug(r *queryResolver, id string) (*models.Drug, error) {
	entity := "drugs"
	whereID := "id = ?"

	record := []*models.Drug{}
	dbRecords := []*dbm.Drug{}

	db := r.ORM.DB.New()
	db = db.Where(whereID, id)
	db = db.Find(&dbRecords)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBDrugToGQLDrug(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record = append(record, rec)
		}
	}

	return record[0], db.Error
}
