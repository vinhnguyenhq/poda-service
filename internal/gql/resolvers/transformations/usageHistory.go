package transformations

import (
	"strconv"

	gql "github.com/vinhnguyenhq/poda-service/internal/gql/models"
	dbm "github.com/vinhnguyenhq/poda-service/internal/orm/models"
)

// DBUsageHistoryToGQLUsageHistory
func DBUsageHistoryToGQLUsageHistory(i *dbm.UsageHistory) (o *gql.UsageHistory, err error) {
	o = &gql.UsageHistory{
		ID:        strconv.FormatUint(uint64(i.ID), 10),
		Quantity:  int(i.Quantity),
		CreatedAt: i.CreatedAt,
	}

	return o, err
}
