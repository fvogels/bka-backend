package filters

import (
	"bass-backend/database/names"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func DocumentNummerBetween(lower int, upper int) Filter {
	return documentNummerFilter{lower: lower, upper: upper}
}

type documentNummerFilter struct {
	lower int
	upper int
}

func (filter documentNummerFilter) Build() squirrel.Sqlizer {
	return squirrel.Expr(
		fmt.Sprintf("%s BETWEEN ? AND ?", names.ColumnDocumentNummer),
		filter.lower,
		filter.upper,
	)
}
