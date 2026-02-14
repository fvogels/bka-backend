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

func InitBoekjaarFilter(storeClause func(squirrel.Sqlizer)) Boekjaar {
	return Boekjaar{
		filter: filter{
			storeClause: storeClause,
		},
	}
}

func (filter Bedrijfsnummer) WithBoekjaar(boekjaar model.BoekJaar) {
	clause := squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.Boekjaar): boekjaar.String(),
	}

	filter.storeClause(clause)
}
