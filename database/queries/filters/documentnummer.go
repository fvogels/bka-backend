package filters

import (
	"bass-backend/database/meta"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type DocumentnummerInterval struct {
	filter
}

func InitDocumentnummerInterval(whereClauses *[]squirrel.Sqlizer) DocumentnummerInterval {
	return DocumentnummerInterval{
		filter: filter{
			whereClauses: whereClauses,
		},
	}
}

func (filter Bedrijfsnummer) WithDocumentNummerBetween(lower string, upper string) {
	clause := squirrel.Expr(
		fmt.Sprintf("%s BETWEEN ? AND ?", meta.DocumentKop.Documentnummer),
		lower,
		upper,
	)

	filter.addWhereClause(clause)
}
