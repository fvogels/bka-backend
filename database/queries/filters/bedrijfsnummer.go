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

func InitBedrijfsnummerFilter(storeClause func(squirrel.Sqlizer)) Bedrijfsnummer {
	return Bedrijfsnummer{
		filter: filter{
			storeClause: storeClause,
		},
	}
}

func (filter Bedrijfsnummer) WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer) {
	clause := squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.Bedrijfsnummer): bedrijfsnummer.String(),
	}

	filter.storeClause(clause)
}
