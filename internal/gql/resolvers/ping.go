package resolvers

import (
	"context"
)

// Ping
func (r *mutationResolver) Ping(ctx context.Context) (bool, error) {
	return ping(r)
}

// ## Helper functions
func ping(r *mutationResolver) (bool, error) {
	return true, nil
}
