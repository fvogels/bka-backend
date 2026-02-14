package filters

import (
	"bass-backend/database/meta"
	"bass-backend/model"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type Bedrijfsnummer struct {
	filter
}

func InitBedrijfsnummerFilter(whereClauses *[]squirrel.Sqlizer) Bedrijfsnummer {
	return Bedrijfsnummer{
		filter: filter{
			whereClauses: whereClauses,
		},
	}
}

func (filter Bedrijfsnummer) WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer) {
	filter.addWhereClause(squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.Bedrijfsnummer): bedrijfsnummer.String(),
	})
}
