package filters

import (
	"bass-backend/database/names"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func DocumentNummerBetween(lower string, upper string) Filter {
	return documentNummerFilter{lower: lower, upper: upper}
}

type documentNummerFilter struct {
	lower string
	upper string
}

func (filter documentNummerFilter) Build() squirrel.Sqlizer {
	return squirrel.Expr(
		fmt.Sprintf("%s BETWEEN ? AND ?", names.ColumnDocumentNummer),
		filter.lower,
		filter.upper,
	)
}
