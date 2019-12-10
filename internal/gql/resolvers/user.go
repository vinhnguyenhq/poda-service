package resolvers

import (
	"context"

	"github.com/vinhnguyenhq/poda-service/internal/gql/models"
	tf "github.com/vinhnguyenhq/poda-service/internal/gql/resolvers/transformations"
	log "github.com/vinhnguyenhq/poda-service/internal/logger"
	dbm "github.com/vinhnguyenhq/poda-service/internal/orm/models"
)

// User
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return user(r, id)
}

// Users
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return users(r)
}

// ## Helper functions
func user(r *queryResolver, id string) (*models.User, error) {
	entity := "users"
	whereID := "id = ?"

	record := []*models.User{}
	dbRecords := []*dbm.User{}

	db := r.ORM.DB.New()
	db = db.Where(whereID, id)
	db = db.Find(&dbRecords)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record = append(record, rec)
		}
	}

	return record[0], db.Error
}

func users(r *queryResolver) ([]*models.User, error) {
	entity := "users"
	record := []*models.User{}
	dbRecords := []*dbm.User{}

	db := r.ORM.DB.New()
	db = db.Find(&dbRecords)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record = append(record, rec)
		}
	}

	return record, db.Error
}
