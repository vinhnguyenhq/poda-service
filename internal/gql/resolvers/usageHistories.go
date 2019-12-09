package resolvers

import (
	"context"

	"github.com/vinhnguyenhq/poda-service/internal/gql/models"
)

// CreateUser creates a record

// Users lists records
func (r *queryResolver) UsageHistories(ctx context.Context, userId *string) (*models.UsageHistories, error) {
	return usageHistoriesList(r, userId)
}

// ## Helper functions
func usageHistoriesList(r *queryResolver, userId *string) (*models.UsageHistories, error) {
	return nil, nil
}
