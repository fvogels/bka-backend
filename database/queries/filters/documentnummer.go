package filters

import (
	"bass-backend/database/meta"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type DocumentnummerInterval struct {
	filter
}

func InitDocumentnummerInterval(storeClause func(squirrel.Sqlizer)) DocumentnummerInterval {
	return DocumentnummerInterval{
		filter: filter{
			storeClause: storeClause,
		},
	}
}

func (filter Bedrijfsnummer) WithDocumentNummerBetween(lower string, upper string) {
	clause := squirrel.Expr(
		fmt.Sprintf("%s BETWEEN ? AND ?", meta.DocumentKop.Documentnummer),
		lower,
		upper,
	)

	filter.storeClause(clause)
}
