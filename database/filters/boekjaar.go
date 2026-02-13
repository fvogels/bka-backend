package filters

import (
	"bass-backend/database/names"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func Boekjaar(boekjaar string) Filter {
	return boekjaarFilter{boekjaar: boekjaar}
}

type boekjaarFilter struct {
	boekjaar string
}

func (filter boekjaarFilter) Build() squirrel.Sqlizer {
	return squirrel.Eq{fmt.Sprintf("%s.%s", names.TableDocumentKop, names.ColumnBoekJaar): filter.boekjaar}
}
