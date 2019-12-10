package transformations

import (
	"strconv"

	gql "github.com/vinhnguyenhq/poda-service/internal/gql/models"
	dbm "github.com/vinhnguyenhq/poda-service/internal/orm/models"
)

// DBUserToGQLUser
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		ID:        strconv.FormatUint(uint64(i.ID), 10),
		Alias:     i.Alias,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
	return o, err
}
