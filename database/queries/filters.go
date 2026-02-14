package queries

import (
	"bass-backend/database/meta"
	"bass-backend/model"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type filter struct {
	whereClauses *[]squirrel.Sqlizer
}

func (filter filter) addWhereClause(clause squirrel.Sqlizer) {
	*filter.whereClauses = append(*filter.whereClauses, clause)
}

type bedrijfsnummerFilter struct {
	filter
}

func initBedrijfsnummerFilter(clauses *[]squirrel.Sqlizer) bedrijfsnummerFilter {
	return bedrijfsnummerFilter{
		filter: filter{
			whereClauses: clauses,
		},
	}
}

func (filter bedrijfsnummerFilter) WithBedrijfsnummer(bedrijfsnummer model.Bedrijfsnummer) {
	filter.addWhereClause(squirrel.Eq{
		fmt.Sprintf("%s.%s", meta.DocumentKop.Table, meta.DocumentKop.BedrijfsNummer): bedrijfsnummer.String(),
	})
}
