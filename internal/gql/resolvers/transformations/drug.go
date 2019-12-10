package transformations

import (
	"strconv"

	gql "github.com/vinhnguyenhq/poda-service/internal/gql/models"
	dbm "github.com/vinhnguyenhq/poda-service/internal/orm/models"
)

// DBUserToGQLUser
func DBDrugToGQLDrug(i *dbm.Drug) (o *gql.Drug, err error) {
	o = &gql.Drug{
		ID:          strconv.FormatUint(uint64(i.ID), 10),
		Name:        i.Name,
		Description: i.Description,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}
	return o, err
}
