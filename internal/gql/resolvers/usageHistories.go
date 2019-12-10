package resolvers

import (
	"context"
	"strconv"

	"github.com/vinhnguyenhq/poda-service/internal/gql/models"
	tf "github.com/vinhnguyenhq/poda-service/internal/gql/resolvers/transformations"
	log "github.com/vinhnguyenhq/poda-service/internal/logger"
	dbm "github.com/vinhnguyenhq/poda-service/internal/orm/models"
)

// Users lists records
func (r *queryResolver) UsageHistories(ctx context.Context, userId *string) ([]*models.UsageHistory, error) {
	return usageHistories(ctx, r, userId)
}

// ## Helper functions
func usageHistories(ctx context.Context, r *queryResolver, userId *string) ([]*models.UsageHistory, error) {
	entity := "usageHistories"
	whereID := "user_id = ?"

	record := []*models.UsageHistory{}
	dbRecords := []*dbm.UsageHistory{}

	db := r.ORM.DB.New()
	if userId != nil {
		db = db.Where(whereID, *userId)
	}

	db = db.Find(&dbRecords)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUsageHistoryToGQLUsageHistory(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			rec.User, err = r.Query().User(ctx, strconv.FormatUint(uint64(dbRec.UserId), 10))
			rec.Drug, err = r.Query().Drug(ctx, strconv.FormatUint(uint64(dbRec.DrugId), 10))

			record = append(record, rec)
		}
	}

	return record, db.Error
}
