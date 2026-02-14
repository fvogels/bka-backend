package filters

import "github.com/Masterminds/squirrel"

type filter struct {
	whereClauses *[]squirrel.Sqlizer
}

func (filter filter) addWhereClause(clause squirrel.Sqlizer) {
	*filter.whereClauses = append(*filter.whereClauses, clause)
}
