package filters

import (
	"bass-backend/database/meta"
	"bass-backend/model"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type Boekjaar struct {
	filter
}

func InitBoekjaarFilter(whereClauses *[]squirrel.Sqlizer) Boekjaar {
	return Boekjaar{
		filter: filter{
			whereClauses: whereClauses,
		},
	}
}

func (filter Bedrijfsnummer) WithBoekjaar(boekjaar model.BoekJaar) {
	filter.addWhereClause(squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.Boekjaar): boekjaar.String(),
	})
}
